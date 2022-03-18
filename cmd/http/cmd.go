package http

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/learninto/goutil"
	"github.com/learninto/goutil/conf"
	"github.com/learninto/goutil/log"

	"github.com/spf13/cobra"
)

var port int
var isInternal bool
var isManage bool

// Cmd run http http
var Cmd = &cobra.Command{
	Use:   "http",
	Short: "Run http",
	Long:  `Run http`,
	Run: func(cmd *cobra.Command, args []string) {
		main()
	},
}

func init() {
	Cmd.Flags().IntVar(&port, "port", 8080, "listen port")
	Cmd.Flags().BoolVar(&isInternal, "internal", false, "internal service")
	Cmd.Flags().BoolVar(&isManage, "manage", false, "manage service")
}

var (
	server *http.Server
	logger = log.Get(context.Background())
)

// 从 http 标准库搬来的
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (net.Conn, error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}
	_ = tc.SetKeepAlive(true)
	_ = tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}

func main() {
	reload := make(chan int, 1)
	stop := make(chan os.Signal, 1)

	conf.OnConfigChange(func() { reload <- 1 })
	conf.WatchConfig()
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)

	startServer()

	for {
		select {
		case <-reload:
			goutil.Reset()
		case sg := <-stop:
			stopServer()
			// 仿 nginx 使用 HUP 信号重载配置
			if sg == syscall.SIGHUP {
				startServer()
			} else {
				goutil.Stop()
				return
			}
		}
	}
}

// startServer
func startServer() {
	logger.Info("start http")

	rand.Seed(int64(time.Now().Nanosecond()))

	mux := http.NewServeMux()

	timeout := 600 * time.Millisecond
	initMux(mux, isInternal)

	if isInternal {
		initInternalMux(mux)

		if d := conf.GetDuration("INTERNAL_API_TIMEOUT"); d > 0 {
			timeout = d * time.Millisecond
		}
	} else {
		if d := conf.GetDuration("OUTER_API_TIMEOUT"); d > 0 {
			timeout = d * time.Millisecond
		}
	}

	panicHandler := goutil.PanicHandler{Handler: mux}
	handler := http.TimeoutHandler(panicHandler, timeout, "timeout")

	if prefix := conf.Get("RPC_PREFIX"); prefix != "" && prefix != "/" {
		handler = http.StripPrefix(prefix, handler)
	}

	http.Handle("/", handler)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/monitor/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	addr := fmt.Sprintf(":%d", port)
	server = &http.Server{IdleTimeout: 60 * time.Second}

	// 配置下发可能会多次触发重启，必须等待 Listen() 调用成功
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		// 本段代码基本搬自 http 标准库
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			panic(err)
		}
		wg.Done()

		err = server.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
		if err != http.ErrServerClosed {
			panic(err)
		}
	}()

	wg.Wait()
}

// stopServer
func stopServer() {
	logger.Info("stop http")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal(err)
	}

	goutil.Reset()
}
