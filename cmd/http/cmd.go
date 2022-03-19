package http

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/learninto/goutil/ctxkit"
	"github.com/learninto/goutil/trace"

	"github.com/prometheus/client_golang/prometheus/promhttp"

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
			Reset()
		case sg := <-stop:
			stopServer()
			// 仿 nginx 使用 HUP 信号重载配置
			if sg == syscall.SIGHUP {
				startServer()
			} else {
				Stop()
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

	panicHandler := PanicHandler{Handler: mux}
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

	Reset()
}

// Reset all utils
func Reset() {
	log.Reset()
}

// Stop all utils
func Stop() {
}

// PanicHandler panic handle
type PanicHandler struct {
	Handler http.Handler
}

// ServeHTTP
func (s PanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	r, span := trace.StartSpanServerHTTP(r, "ServeHTTP") // 开始链路
	defer func() {
		if rec := recover(); rec != nil {
			ctx := r.Context()
			ctx = ctxkit.WithTraceID(ctx, trace.GetTraceID(ctx))
			log.Get(ctx).Error(rec, string(debug.Stack()))
		}
		span.Finish()
	}()

	origin := r.Header.Get("Origin")
	suffix := conf.Get("CORS_ORIGIN_SUFFIX")

	if origin != "" && suffix != "" && strings.HasSuffix(origin, suffix) {
		w.Header().Add("Access-Control-Allow-Origin", origin)
		w.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Origin,No-Cache,X-Requested-With,If-Modified-Since,Pragma,Last-Modified,Cache-Control,Expires,Content-Type,Access-Control-Allow-Credentials,DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Cache-Webcdn,Content-Length")
	}

	if r.Method == http.MethodOptions {
		return
	}

	s.Handler.ServeHTTP(w, r)
}
