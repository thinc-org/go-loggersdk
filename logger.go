package loggersdk

import (
	"github.com/thinc-org/go-loggersdk/pkg/sentry"
	"go.uber.org/zap"
)

type Logger struct {
	serviceName string
	*zap.SugaredLogger
	sentryService sentry.Sentry
}

func NewService(serviceName string, dsn string) (*Logger, error) {
	lg, _ := zap.NewProduction()
	sugar := lg.Sugar()

	sentrySrv, err := sentry.NewService(dsn)
	if err != nil {
		return nil, err
	}

	return &Logger{serviceName, sugar, sentrySrv}, nil
}
func (s *Logger) SetName(name string) {
	s.serviceName = name
}

func (s *Logger) Info(msg string, options ...interface{}) {
	if len(options) > 0 {
		s.SugaredLogger.Infow(msg, "service_name", s.serviceName, options)
		return
	}

	s.SugaredLogger.Infow(msg, "service_name", s.serviceName)
}

func (s *Logger) Error(msg string, err error, options ...interface{}) {
	s.sentryService.CaptureException(err)
	if len(options) > 0 {
		s.SugaredLogger.Errorw(msg, "error", zap.Error(err), "service_name", s.serviceName, options)
		return
	}

	s.SugaredLogger.Errorw(msg, "error", zap.Error(err), "service_name", s.serviceName)
}

func (s *Logger) Debug(msg string, options ...interface{}) {
	if len(options) > 0 {
		s.SugaredLogger.Debugw(msg, "service_name", s.serviceName, options)
		return
	}

	s.SugaredLogger.Debugw(msg, "service_name", s.serviceName)
}

func (s *Logger) Fatal(msg string, options ...interface{}) {
	if len(options) > 0 {
		s.SugaredLogger.Fatalw(msg, "service_name", s.serviceName, options)
		return
	}

	s.SugaredLogger.Fatalw(msg, "service_name", s.serviceName)
}
