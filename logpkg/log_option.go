package logpkg

// LogOption 日志配置项
type LogOption func(zl *zapLogger)

// NewLogger 初始化日志记录器
func NewLogger(options ...LogOption) error {
	var err error

	logger, err = newZapLogger(options...)
	if err != nil {
		return err
	}
	return nil
}

// WithName 日志输出时的应用名
func WithName(n string) LogOption {
	return func(zl *zapLogger) {
		zl.name = n
	}
}

// WithWriters 日志输出流, file: 只输出到文件; stdout: 只输出到标准输出流; file,stdout: 同时输出到文件和标准输出流
func WithWriters(w string) LogOption {
	return func(zl *zapLogger) {
		zl.writers = w
	}
}

// WithLevel 日志级别: DEBUG, INFO, WARN, ERROR, FATAL
func WithLevel(l string) LogOption {
	return func(zl *zapLogger) {
		zl.loggerLevel = l
	}
}

// WithFile 日志文件名: name.logger
func WithFile(f string) LogOption {
	return func(zl *zapLogger) {
		zl.loggerFile = f
	}
}

// WithWarnFile 警告日志文件名: name_warn.logger
func WithWarnFile(f string) LogOption {
	return func(zl *zapLogger) {
		zl.loggerWarnFile = f
	}
}

// WithErrorFile 错误日志名: name_error.logger
func WithErrorFile(f string) LogOption {
	return func(zl *zapLogger) {
		zl.loggerErrorFile = f
	}
}

// WithFormatText 日志输出格式, json, plaintext
func WithFormatText(b bool) LogOption {
	return func(zl *zapLogger) {
		zl.logFormatText = b
	}
}

// WithRollingPolicy 日志切割方式, daily: 每24小时切割; hourly: 每小时切割
func WithRollingPolicy(r string) LogOption {
	return func(zl *zapLogger) {
		zl.logRollingPolicy = r
	}
}

// WithRotateDate 日志转存时间
func WithRotateDate(d int) LogOption {
	return func(zl *zapLogger) {
		zl.logRotateDate = d
	}
}

// WithRotateSize 日志转存大小
func WithRotateSize(s int) LogOption {
	return func(zl *zapLogger) {
		zl.logRotateSize = s
	}
}

// WithBackUpCount 日志数量达到一定量后进行压缩备份
func WithBackUpCount(n uint) LogOption {
	return func(zl *zapLogger) {
		zl.logBackupCount = n
	}
}
