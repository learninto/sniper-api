package hooks

import (
	"github.com/learninto/goutil/conf"
	"github.com/prometheus/client_golang/prometheus"
)

var defBuckets = []float64{.005, .01, .025, .05, .1, .25, .5, 1}
var rpcDurations = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Namespace: conf.AppID,
	Subsystem: "rpc",
	Name:      "server_durations_seconds",
	Help:      "RPC latency distributions",
	Buckets:   defBuckets,
}, []string{"path", "code"})

var LogTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Namespace:   conf.AppID,
	Name:        "log_total",
	Help:        "log total",
	ConstLabels: map[string]string{"app": conf.AppID},
}, []string{"code"})

var JobTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Namespace:   conf.AppID,
	Name:        "job_total",
	Help:        "job total",
	ConstLabels: map[string]string{"app": conf.AppID},
}, []string{"code"})

var MQDurationsSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Namespace:   conf.AppID,
	Name:        "mq_durations_seconds",
	Help:        "Databus latency distributions",
	Buckets:     defBuckets,
	ConstLabels: map[string]string{"app": conf.AppID},
}, []string{"name", "role"})

var NetPoolHits = prometheus.NewCounterVec(prometheus.CounterOpts{
	Namespace:   conf.AppID,
	Name:        "net_pool_hits",
	Help:        "net pool hits",
	ConstLabels: map[string]string{"app": conf.AppID},
}, []string{"name", "type"})

var NetPoolMisses = prometheus.NewCounterVec(prometheus.CounterOpts{
	Namespace:   conf.AppID,
	Name:        "net_pool_misses",
	Help:        "net pool misses",
	ConstLabels: map[string]string{"app": conf.AppID},
}, []string{"name", "type"})

var NetPoolTimeouts = prometheus.NewCounterVec(prometheus.CounterOpts{
	Namespace:   conf.AppID,
	Name:        "net_pool_timeouts",
	Help:        "net pool timeouts",
	ConstLabels: map[string]string{"app": conf.AppID},
}, []string{"name", "type"})

var NetPoolStale = prometheus.NewCounterVec(prometheus.CounterOpts{
	Namespace:   conf.AppID,
	Name:        "net_pool_stale",
	Help:        "net pool stale",
	ConstLabels: map[string]string{"app": conf.AppID},
}, []string{"name", "type"})

var NetPoolTotal = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace:   conf.AppID,
	Name:        "net_pool_total",
	Help:        "net pool total",
	ConstLabels: map[string]string{"app": conf.AppID},
}, []string{"name", "type"})

var NetPoolIdle = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace:   conf.AppID,
	Name:        "net_pool_idle",
	Help:        "net pool idle",
	ConstLabels: map[string]string{"app": conf.AppID},
}, []string{"name", "type"})

var RPCDurationsSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Namespace:   conf.AppID,
	Name:        "rpc_durations_seconds",
	Help:        "RPC latency distributions",
	Buckets:     defBuckets,
	ConstLabels: map[string]string{"app": conf.AppID},
}, []string{"path", "code"})

func init() {
	prometheus.MustRegister(rpcDurations)
	prometheus.MustRegister(LogTotal)
	prometheus.MustRegister(JobTotal)
	prometheus.MustRegister(MQDurationsSeconds)
	prometheus.MustRegister(NetPoolHits)
	prometheus.MustRegister(NetPoolMisses)
	prometheus.MustRegister(NetPoolTimeouts)
	prometheus.MustRegister(NetPoolStale)
	prometheus.MustRegister(NetPoolTotal)
	prometheus.MustRegister(NetPoolIdle)
	prometheus.MustRegister(RPCDurationsSeconds)
}
