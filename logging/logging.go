package logging

import (
	"github.com/Sirupsen/logrus"
	"github.com/go-errors/errors"
	"gopkg.in/natefinch/lumberjack.v2"
)

type WrappedLogger struct {
	logger *logrus.Logger
}

var AppLogger WrappedLogger

func Create(logPath string) WrappedLogger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})

	logger.SetOutput(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})

	return WrappedLogger{logger: logger}
}

func AppLoggerInit(logPath string) {
	AppLogger = Create(logPath)
}
func (w WrappedLogger) Info(msg string) {
	w.logger.Info(msg)
}
func (w WrappedLogger) Error(e error) {
	errorWithStack, ok := e.(*errors.Error)
	if ok {
		w.logger.Error(errorWithStack.ErrorStack())
	} else {
		w.logger.Error(e.Error())
	}
}
