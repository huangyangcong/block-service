package schedule

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func NewBoxPrice(c *cron.Cron) *cron.Cron {
	c.AddFunc("@every 1s", func() {
		fmt.Errorf("aaaa")
	})
	return c
}
