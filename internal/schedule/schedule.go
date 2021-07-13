package schedule

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/robfig/cron/v3"
)

// ProviderSet is schedule providers.
var ProviderSet = wire.NewSet(NewScheduleServer, NewRouter, NewBoxPrice)

type Server struct {
	log      log.Logger
	schedule *cron.Cron
}
type Router struct {
	boxPrice BoxPrice
}

func NewRouter(price BoxPrice) *Router {
	return &Router{price}
}

func NewScheduleServer(logger log.Logger) *Server {
	l := Logger{log.NewHelper(logger)}
	schedule := cron.New(
		cron.WithLogger(l),
		cron.WithChain(cron.Recover(l)),
	)
	return &Server{log: logger, schedule: schedule}
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

func (s Server) Start(c context.Context) error {
	s.schedule.Start()
	return nil
}
func (s Server) Stop(c context.Context) error {
	s.schedule.Stop()
	return nil
}
