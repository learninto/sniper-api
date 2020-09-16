package init

import (
	"github.com/learninto/goutil/conf"
	"github.com/learninto/goutil/metrics"

	_ "github.com/learninto/goutil/init"
)

func init() {
	metrics.InitMetrics(conf.AppID)
}
