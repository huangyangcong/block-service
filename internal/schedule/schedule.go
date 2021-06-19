package schedule

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/robfig/cron/v3"
)

// ProviderSet is schedule providers.
var ProviderSet = wire.NewSet(NewSchedule, NewScheduleRoutes, NewBoxPrice)

type Routes struct {
}

func NewScheduleRoutes(price BoxPrice) *Routes {
	return &Routes{}
}

type Logger struct {
	logger *log.Helper
}

func (l Logger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Infof(msg, keysAndValues...)
}

// Error logs an error condition.
func (l Logger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.logger.Errorf(msg, keysAndValues...)
}

func NewSchedule(logger log.Logger) *cron.Cron {
	l := Logger{log.NewHelper(logger)}
	c := cron.New(
		cron.WithLogger(l),
		cron.WithChain(cron.Recover(l)),
	)
	c.Start()
	return c
}
