package driver

import (
	"fmt"
	"github.com/gongxulei/go_kit/logger/constant"
	"github.com/gongxulei/go_kit/logger/origin"
	"os"
	"runtime"
	"time"
)

// 2021-11-18 23:59:59 Info Test.go:29 this is a info log
type FileLogger struct {
	level             uint8
	logPath           string
	logName           string
	logFileHandler    *os.File
	outputConsole     bool
	LogMessageChannel chan *LogMessage
	//logErrorFileHandler *os.File
}

// 异步写入日志 1、创建一个异步的channel，2、

type LogMessage struct {
	messageStr string
	nowTimeStr string
	level      string
	fileName   string
	lineNo     int
}

func NewFileLogger(level uint8, logPath string, logName string) origin.LoggerInterface {
	var channelLen = 50000
	logger := &FileLogger{
		level:             level,
		logPath:           logPath,
		logName:           logName,
		outputConsole:     true, //默认输出console控制台日志
		LogMessageChannel: make(chan *LogMessage, channelLen),
	}
	logger.init()
	return logger
}

func (f *FileLogger) init() {
	fileName := fmt.Sprintf("%s/%s_info.log", f.logPath, f.logName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed, err:%v", fileName, err.Error()))
	}
	f.logFileHandler = file

	// 错误日志和fatal日志文件
	//fileName = fmt.Sprintf("%s/%s_error.log", f.logPath, f.logName)
	//file, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	//if err != nil {
	//	panic(fmt.Sprintf("open file %s failed, err:%v", fileName, err.Error()))
	//}
	//f.logErrorFileHandler = file

	go func() {
		for message := range f.LogMessageChannel {
			fmt.Fprintln(f.logFileHandler,
				message.nowTimeStr,
				message.level,
				fmt.Sprintf("%s:%d", message.fileName, message.lineNo),
				message.messageStr,
			)
			fmt.Fprintln(os.Stdout,
				message.nowTimeStr,
				message.level,
				fmt.Sprintf("%s:%d", message.fileName, message.lineNo),
				message.messageStr)
		}
	}()

}

func (f *FileLogger) SetLevel(level uint8) {
	if level < constant.LogLevelDebug || level > constant.LogLevelFatal {
		level = constant.LogLevelDebug
	}
	f.level = level
}

func (f *FileLogger) LogDebug(format string, args ...interface{}) {
	f.writeLog(constant.LogLevelDebug, format, args)
}

func (f *FileLogger) LogTrace(format string, args ...interface{}) {
	f.writeLog(constant.LogLevelTrace, format, args)
}

func (f *FileLogger) LogInfo(format string, args ...interface{}) {
	f.writeLog(constant.LogLevelInfo, format, args)
}
func (f *FileLogger) LogWaring(format string, args ...interface{}) {
	f.writeLog(constant.LogLevelWaring, format, args)
}
func (f *FileLogger) LogError(format string, args ...interface{}) {
	f.writeLog(constant.LogLevelError, format, args)
}
func (f *FileLogger) LogFatal(format string, args ...interface{}) {
	f.writeLog(constant.LogLevelFatal, format, args)
}

func (f *FileLogger) Log(format string, args ...interface{}) {
	f.writeLog(f.level, format, args)
}

func (f *FileLogger) Close() {
	_ = f.logFileHandler.Close()
}

func (f *FileLogger) writeLog(level uint8, format string, args ...interface{}) {
	now := time.Now()
	if level < f.level {
		return
	}
	//file := f.logFileHandler
	//if level >= constant.LogLevelError {
	//	file = f.logErrorFileHandler
	//}

	fileName, _, lineNo := getLineInfo()

	// 将数据写入channel
	select {
	case f.LogMessageChannel <- &LogMessage{
		messageStr: fmt.Sprintf(format, args...),
		nowTimeStr: now.Format("2006-01-02 15:04:05.999"),
		level:      constant.GetLogLevel(level),
		fileName:   fileName,
		lineNo:     lineNo,
	}:
	default:
	}

	//logFormatSlice := []string{
	//	now.Format("2006-01-02 15:04:05.999"),
	//	constant.GetLogLevel(level),
	//	fmt.Sprintf("%s:%d", fileName, lineNo),
	//	funcName,
	//	format,
	//	"\n",
	//}
	//format = strings.Join(logFormatSlice, " ")
	//if f.outputConsole {
	//	fmt.Printf(format, args...)
	//}
	//fmt.Fprintf(file, format, args...)
}

func getLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}
