package schedule

import (
	"github.com/google/wire"
	"github.com/robfig/cron/v3"
	"time"
)

// ProviderSet is schedule providers.
var ProviderSet = wire.NewSet(NewSchedule)

type Schedule struct {
	timeZone string
	cronImpl *cron.Cron
	jobs     Job
}
type Job struct {
	runAt string
	cmd   func()
}

func NewSchedule() *Schedule {
	wire.Build(wire.Struct(new(Schedule), "jobs"))
	return &Schedule{}
}
func (c *Schedule) Run() error {
	if err := c.init(); err != nil {
		return err
	}

	c.cronImpl.Start()
	return nil
}

func (c *Schedule) init() error {
	if c.cronImpl != nil {
		c.cronImpl.Stop()
	}
	if c.timeZone == "none" {
		c.cronImpl = cron.New()
	} else {
		tz, err := time.LoadLocation(c.timeZone)
		if err != nil {
			return err
		}

		c.cronImpl = cron.New(
			cron.WithLocation(tz),
		)
	}

	for _, job := range c.jobs {
		c.cronImpl.AddFunc(job.runAt, job.cmd)
	}
	return nil
}

func (c *Schedule) Stop() {
	c.cronImpl.Stop()
}
