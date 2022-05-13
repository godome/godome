package logger

var logger = newStandardLogger()

func GetLogger() Logger {
	return logger
}

func SetLogger(newLoggger Logger) {
	logger = newLoggger
}

type Logger interface {
	Info(msg string, a ...interface{})
	Debug(msg string, a ...interface{})
	Warn(msg string, a ...interface{})
	Fatal(msg string, a ...interface{})
	Error(msg string, a ...interface{})
	Panic(msg string, a ...interface{})
}

func Info(msg string, args ...interface{}) {
	logger.Info(msg, args...)
}

func Debug(msg string, args ...interface{}) {
	logger.Debug(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	logger.Warn(msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	logger.Fatal(msg, args...)
}

func Error(msg string, args ...interface{}) {
	logger.Error(msg, args...)
}

func Panic(msg string, args ...interface{}) {
	logger.Panic(msg, args...)
}
