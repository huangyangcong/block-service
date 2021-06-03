// +build wireinject

package schedule

import (
	"github.com/google/wire"
	"github.com/robfig/cron/v3"
)

func InitSchedule() (*cron.Cron, error) {
	panic(wire.Build(NewBoxPrice, NewSchedule))
}
func NewSchedule() *cron.Cron {
	c := cron.New()
	c.Start()
	defer c.Stop()
	return c
}
