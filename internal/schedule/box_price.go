package schedule

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
)

type BoxPrice struct {
}

func NewBoxPrice(c *cron.Cron, logger log.Logger) BoxPrice {
	c.AddFunc("@every 1s", func() {
		logger.Log(log.LevelDebug, "aaaa")
	})
	return BoxPrice{}
}
