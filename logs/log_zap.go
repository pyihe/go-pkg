package logs

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/pyihe/go-pkg/nets"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// writerStdOut 标准输出流
	writerStdOut = "stdout"

	// writerFile  文件输出流
	writerFile = "file"
)

const (
	// RotateTimeDaily 默认按天切割
	RotateTimeDaily = "daily"
	// RotateTimeHourly 按小时切割
	RotateTimeHourly = "hourly"
)

type zapLogger struct {
	name             string // 日志对应的应用名
	writers          string // 日志输出流
	loggerLevel      string // 日志级别
	loggerFile       string // 日志文件
	loggerWarnFile   string // 警告日志文件
	loggerErrorFile  string // 错误日志文件
	logFormatText    bool   // 日志输出格式
	logRollingPolicy string // 日志文件切割策略
	logRotateDate    int    // 日志文件转存时间
	logRotateSize    int    // 日志转存大小
	logBackupCount   uint   // 日志文件达到多少个时进行压缩备份

	sugaredLogger *zap.SugaredLogger
}

func newZapLogger(opts ...LogOption) (Logger, error) {
	var (
		zlogger = &zapLogger{}
		cores   []zapcore.Core
		options []zap.Option
	)

	for _, opt := range opts {
		opt(zlogger)
	}

	encoder := getJSONEncoder()

	op := zap.Fields(zap.String("ip", nets.GetLocalIP()), zap.String("app", zlogger.name))
	options = append(options, op)

	allLevel := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		return lv <= zapcore.FatalLevel
	})

	writers := strings.Split(zlogger.writers, ",")
	for _, w := range writers {
		if w == writerStdOut {
			c := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
			cores = append(cores, c)
		}
		if w == writerFile {
			infoWriter := zlogger.getLogWriterWithTime(zapcore.InfoLevel)
			warnWriter := zlogger.getLogWriterWithTime(zapcore.WarnLevel)
			errorWriter := zlogger.getLogWriterWithTime(zapcore.ErrorLevel)

			infoLevel := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
				return lv <= zapcore.InfoLevel
			})
			warnLevel := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
				stackTrace := zap.AddStacktrace(zapcore.WarnLevel)
				options = append(options, stackTrace)
				return lv == zapcore.WarnLevel
			})
			errorLevel := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
				stackTrace := zap.AddStacktrace(zapcore.ErrorLevel)
				options = append(options, stackTrace)
				return lv >= zapcore.ErrorLevel
			})

			core := zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel)
			cores = append(cores, core)
			core = zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel)
			cores = append(cores, core)
			core = zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel)
			cores = append(cores, core)
		}
		if w != writerFile && w != writerStdOut {
			core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
			cores = append(cores, core)
			allWriter := zlogger.getLogWriterWithTime(zapcore.InfoLevel)
			core = zapcore.NewCore(encoder, zapcore.AddSync(allWriter), allLevel)
			cores = append(cores, core)
		}
	}

	combinedCore := zapcore.NewTee(cores...)

	// 开启开发模式， 堆栈跟踪
	caller := zap.AddCaller()
	options = append(options, caller)

	//开启文件及行号
	development := zap.Development()
	options = append(options, development)

	//跳过文件调用层数
	addCallerSkip := zap.AddCallerSkip(2)
	options = append(options, addCallerSkip)

	zlogger.sugaredLogger = zap.New(combinedCore, options...).Sugar()

	return zlogger, nil
}

// getJSONEncoder
func getJSONEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "app",
		CallerKey:      "file",
		StacktraceKey:  "trace",
		LineEnding:     "",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     nil,
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getLogWriterWithTime 按时间对日志进行切割(默认24小时切割一次)
func (z *zapLogger) getLogWriterWithTime(level zapcore.Level) io.Writer {
	var filename string
	switch level {
	case zapcore.WarnLevel:
		filename = z.loggerWarnFile
	case zapcore.ErrorLevel:
		filename = z.loggerErrorFile
	default:
		filename = z.loggerFile
	}

	//默认24小时切割一次
	rotateDuration := time.Hour * 24
	if z.logRollingPolicy == RotateTimeHourly {
		rotateDuration = time.Hour
	}
	hook, err := rotatelogs.New(
		filename+".%Y%m%d%H",                           //保存的文件名，后缀为时间
		rotatelogs.WithLinkName(filename),              //生成软联，指向最新的日志文件
		rotatelogs.WithRotationCount(z.logBackupCount), //日志文件最多保存多少份
		rotatelogs.WithRotationTime(rotateDuration),    //切割的时间间隔
	)
	if err != nil {
		panic(err)
	}
	return hook
}

func (z *zapLogger) Debug(args ...interface{}) {
	z.sugaredLogger.Debug(args...)
}

func (z *zapLogger) Debugf(format string, args ...interface{}) {
	z.sugaredLogger.Debugf(format, args...)
}

func (z *zapLogger) Info(args ...interface{}) {
	z.sugaredLogger.Info(args...)
}

func (z *zapLogger) Infof(format string, args ...interface{}) {
	z.sugaredLogger.Infof(format, args...)
}

func (z *zapLogger) Warn(args ...interface{}) {
	z.sugaredLogger.Warn(args...)
}

func (z *zapLogger) Warnf(format string, args ...interface{}) {
	z.sugaredLogger.Warnf(format, args...)
}

func (z *zapLogger) Error(args ...interface{}) {
	z.sugaredLogger.Error(args...)
}

func (z *zapLogger) Errorf(format string, args ...interface{}) {
	z.sugaredLogger.Errorf(format, args...)
}

func (z *zapLogger) Fatal(args ...interface{}) {
	z.sugaredLogger.Fatal(args...)
}

func (z *zapLogger) Fatalf(format string, args ...interface{}) {
	z.sugaredLogger.Fatalf(format, args...)
}

func (z *zapLogger) Panicf(format string, args ...interface{}) {
	z.sugaredLogger.Panicf(format, args...)
}

func (z *zapLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k, v)
	}
	newLogger := z.sugaredLogger.With(f...)
	return &zapLogger{sugaredLogger: newLogger}
}
