package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	log *zap.SugaredLogger
}

func NewLogger() *Logger {
	zapLogger, _ := zap.NewProduction()

	sugar := zapLogger.Sugar()

	return &Logger{
		log: sugar,
	}
}

func (l *Logger) Error(msg string, err error) {
	l.log.Errorw(msg, "error", zap.Error(err))
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	l.log.Infof(msg, args...)
}
func (l *Logger) Warn(args ...interface{}) {
	l.log.Warnln(args...)
}
