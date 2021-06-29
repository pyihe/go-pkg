package logs

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
)

const (
	colorRed    = 31
	colorYellow = 33
	colorBlue   = 34

	levelT = "[T] "
	levelE = "[E] "
	levelW = "[W] "
	levelF = "[F] "

	defaultFileSize = 60 * 1024 * 1024
	minFileSize     = 1 * 1024 * 1024
	defaultLogDir   = "logs"
	defaultLogName  = "default.logs"
)

const (
	logTypeStd logType = iota + 1
	logTypeFile
)

type (
	logType int

	logOption func(log *myLog)

	myLog struct {
		sync.Once
		sync.Mutex
		outs     map[logType]io.Writer //writer集合
		file     *os.File              //文件句柄
		fileName string                //日志名
		dir      string                //日志存放路径
		size     int64                 //单个日志文件的大小限制
	}
)

var (
	defaultLogger = &myLog{}
)

func (m *myLog) init() {
	if m.dir == "" {
		m.dir = defaultLogDir
	}
	if m.fileName == "" {
		m.fileName = defaultLogName
	}
	if m.size == 0 {
		m.size = defaultFileSize
	} else {
		if m.size < minFileSize {
			panic(fmt.Sprintf("invalid size: %d", m.size))
		}
	}

	if m.outs == nil {
		m.outs = make(map[logType]io.Writer)
	}
	if !isExist(m.dir) {
		if err := os.Mkdir(m.dir, 0777); err != nil {
			panic(err)
		}
	}
	name := path.Join(m.dir, m.fileName)
	file, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}

	m.file = file
	m.outs[logTypeStd] = os.Stdout
	m.outs[logTypeFile] = file
}

func (m *myLog) checkLogSize() {
	if m.file == nil {
		return
	}
	m.Lock()
	defer m.Unlock()
	fileInfo, err := m.file.Stat()
	if err != nil {
		panic(err)
	}
	if m.size > fileInfo.Size() {
		return
	}
	//需要分割
	newName := path.Join(m.dir, time.Now().Format("2006_01_02_15:04:03")+".logs")
	name := path.Join(m.dir, m.fileName)

	err = os.Rename(name, newName)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}

	m.file.Close()
	m.file = file
	m.outs[logTypeFile] = file
	return
}

func (m *myLog) writeLog() {

}

func (m *myLog) write(msg, timeStr, fileName, prefix string, line, color int) {
	m.checkLogSize()

	for k, wr := range m.outs {
		if k == logTypeStd {
			fmt.Fprintf(wr, "%c[%dm%s[%s %s:%d] %s%c[0m\n", 0x1B, color, prefix, timeStr, fileName, line, msg, 0x1B)
		} else {
			fmt.Fprintf(wr, "%s[%s %s:%d] %s\n", prefix, timeStr, fileName, line, msg)
		}
	}
}

func WithSize(size int64) logOption {
	return func(log *myLog) {
		log.size = size
	}
}

func WithLogDir(dir string) logOption {
	return func(log *myLog) {
		log.dir = dir
	}
}

func WithFileName(name string) logOption {
	return func(log *myLog) {
		log.fileName = name
	}
}

func InitLogger(args ...logOption) {
	defaultLogger.Do(func() {
		for _, af := range args {
			af(defaultLogger)
		}
		defaultLogger.init()
	})
}

////Info
func T(format string, v ...interface{}) {
	timeStr, fileName, line := getPrefixInfo()
	defaultLogger.write(fmt.Sprintf(format, v...), timeStr, fileName, levelT, line, colorBlue)
}

//
////Error
func E(format string, v ...interface{}) {
	timeStr, fileName, line := getPrefixInfo()
	defaultLogger.write(fmt.Sprintf(format, v...), timeStr, fileName, levelE, line, colorYellow)
}

//Warn
func W(format string, v ...interface{}) {
	timeStr, fileName, line := getPrefixInfo()
	defaultLogger.write(fmt.Sprintf(format, v...), timeStr, fileName, levelW, line, colorRed)
}

//Fatal
func F(format string, v ...interface{}) {
	timeStr, fileName, line := getPrefixInfo()
	defaultLogger.write(fmt.Sprintf(format, v...), timeStr, fileName, levelF, line, colorRed)
	os.Exit(-1)
}

func getPrefixInfo() (timeStr, fileName string, line int) {
	_, file, line, _ := runtime.Caller(1)
	fileName = shortFileName(file)
	timeStr = time.Now().Format("2006-01-02 15:04:05.0000")
	return timeStr, fileName, line
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func shortFileName(file string) string {
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	return short
}
