package cron

import (
	"github.com/learninto/goutil/conf"
	"github.com/prometheus/client_golang/prometheus"
)

var defBuckets = []float64{.005, .01, .025, .05, .1, .25, .5, 1}
var JobTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Namespace:   conf.AppID,
	Name:        "job_total",
	Help:        "job total",
	ConstLabels: map[string]string{"app": conf.AppID},
}, []string{"code"})

func init() {
	prometheus.MustRegister(JobTotal)
}
