package constant


// 日志级别
const (
	LogLevelDebug uint8 = iota
	LogLevelTrace
	LogLevelInfo
	LogLevelWaring
	LogLevelError
	LogLevelFatal
)

func GetLogLevel(level uint8) string {
	if level == LogLevelDebug {
		return "Debug"
	}
	if level == LogLevelTrace {
		return "Trace"
	}
	if level == LogLevelInfo {
		return "Info"
	}
	if level == LogLevelWaring {
		return "Waring"
	}
	if level == LogLevelError {
		return "Error"
	}
	if level == LogLevelFatal {
		return "Fatal"
	}
	return "Info"
}