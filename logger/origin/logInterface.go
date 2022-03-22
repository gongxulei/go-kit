package origin

type LoggerInterface interface {
	SetLevel(level uint8)
	LogDebug(format string, args ...interface{})
	LogTrace(format string, args ...interface{})
	LogInfo(format string, args ...interface{})
	LogWaring(format string, args ...interface{})
	LogError(format string, args ...interface{})
	LogFatal(format string, args ...interface{})
	Log(format string, args ...interface{})
	Close()
}
