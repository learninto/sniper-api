package init

import (
	"kingstar-go/sniper/conf"
	"kingstar-go/sniper/metrics"

	_ "kingstar-go/sniper/init"
)

func init() {
	metrics.InitMetrics(conf.AppID)
}
