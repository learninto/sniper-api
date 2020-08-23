package init

import (
	"github.com/learninto/sniper-api/utils/conf"
	"github.com/learninto/sniper-api/utils/metrics"

	_ "github.com/learninto/sniper-api/utils/init"
)

func init() {
	metrics.InitMetrics(conf.AppID)
}
