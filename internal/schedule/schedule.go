package schedule

import (
	"time"

	"github.com/google/wire"
	"github.com/robfig/cron/v3"
	robfigCron "github.com/robfig/cron/v3"
)

// ProviderSet is schedule providers.
var ProviderSet = wire.NewSet(NewSchedule)

func NewSchedule() *cron.Cron {
	//cron
	s := NewCron("none")
	if err := s.Start(); err != nil {
		panic(err)
	}
	s.Push("@every 1s", NewBoxPrice())

	defer s.Stop()
	return s.cronImpl
}

type Cron struct {
	timeZone string
	cronImpl *robfigCron.Cron
	jobs     map[string]func()
}

func NewCron(timeZone string) *Cron {
	return &Cron{
		timeZone: timeZone,
		jobs:     make(map[string]func()),
	}
}

func (c *Cron) Push(runat string, job func()) {
	c.jobs[runat] = job
}

func (c *Cron) Start() error {
	if err := c.init(); err != nil {
		return err
	}

	c.cronImpl.Start()
	return nil
}

func (c *Cron) init() error {
	if c.cronImpl != nil {
		c.cronImpl.Stop()
	}
	if c.timeZone == "none" {
		c.cronImpl = robfigCron.New()
	} else {
		tz, err := time.LoadLocation(c.timeZone)
		if err != nil {
			return err
		}

		c.cronImpl = robfigCron.New(
			robfigCron.WithLocation(tz),
		)
	}

	for runat, job := range c.jobs {
		c.cronImpl.AddFunc(runat, job)
	}
	return nil
}

func (c *Cron) Stop() {
	c.cronImpl.Stop()
}