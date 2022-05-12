package logger

import (
	"fmt"
	"log"
	"os"
)

type standardLogger struct {
	logger *log.Logger
}

func newStandardLogger() Logger {
	return &standardLogger{
		logger: log.New(os.Stdout, "", 0),
	}
}

func (r *standardLogger) Info(msg string, args ...interface{}) {
	r.logger.Printf(msg, args...)
}

func (r *standardLogger) Debug(msg string, args ...interface{}) {
	if os.Getenv("DEBUG") == "1" || os.Getenv("DEBUG") == "true" {
		r.logger.Printf(msg, args...)
	}
}

func (r *standardLogger) Warn(msg string, args ...interface{}) {
	r.logger.Printf(msg, args...)
}

func (r *standardLogger) Fatal(msg string, args ...interface{}) {
	r.logger.Fatal(fmt.Sprintf(msg, args...))
}

func (r *standardLogger) Error(msg string, args ...interface{}) {
	r.logger.Fatal(fmt.Sprintf(msg, args...))
}

func (r *standardLogger) Panic(msg string, args ...interface{}) {
	r.logger.Panic(fmt.Sprintf(msg, args...))
}
