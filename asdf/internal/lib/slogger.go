package lib

import (
	"io"
	"log"
	"log/slog"
	"os"
)

type slogLogger struct {
	log *slog.Logger
}

var (
	_      ILogger = (*slogLogger)(nil)
	logger *slogLogger
)

func newLogger(out io.Writer) *slogLogger {
	if logger != nil {
		return logger
	}

	if out == nil {
		out = os.Stdout
	}

	sl := slog.New(slog.NewJSONHandler(out, nil))
	logger = &slogLogger{
		log: sl,
	}

	return logger
}

func (sl *slogLogger) Debug(msg string, keysAndValues ...interface{}) {
	sl.log.Debug(msg, keysAndValues...)
}

func (sl *slogLogger) Info(msg string, keysAndValues ...interface{}) {
	sl.log.Info(msg, keysAndValues...)
}

func (sl *slogLogger) Warn(msg string, keysAndValues ...interface{}) {
	sl.log.Warn(msg, keysAndValues...)
}

func (sl *slogLogger) Error(msg string, keysAndValues ...interface{}) {
	sl.log.Error(msg, keysAndValues...)
}

func (sl *slogLogger) Fatal(msg string, keysAndValues ...interface{}) {
	sl.log.Error(msg, keysAndValues...)
	log.Fatal(msg)
}

func (sl *slogLogger) Handler() slog.Handler {
	return sl.log.Handler()
}
