package logs

var (
	logger Logger
)

type Fields map[string]interface{}

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	WithFields(fields Fields) Logger
}

// Debug
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Info
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Warn
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Error
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Fatal
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Panicf
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

// WithFields
func WithFields(fields Fields) Logger {
	return logger.WithFields(fields)
}
