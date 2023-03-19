package lightlog

type Logger struct {
	*LoggerCore
}

type LoggerOptions struct {
	name     string
	level    Level
	filepath string
}

func NewLogger(options *LoggerOptions) *Logger {
	if options.name == "" {
		options.name = "defaultLogger"
	}

	logger := &Logger{
		NewLoggerCore(options.name, options.level),
	}
	logger.AddTransport("defaultConsoleTransport", NewConsoleTransport("defaultConsoleTransport", options.level))

	if options.filepath == "" {
		logger.Warn("`options.filepath` cannot be empty")
	} else {
		logger.AddTransport("defaultFileTransport", NewFileTransport("defaultFileTransport", options.level, options.filepath))
	}

	return logger
}

type ConsoleLogger struct {
	*LoggerCore
}

func NewConsoleLogger(name string, level Level) *ConsoleLogger {
	logger := &ConsoleLogger{NewLoggerCore(name, level)}
	logger.AddTransport("defaultConsoleTransport", NewConsoleTransport("defaultConsoleTransport", level))

	return logger
}

type FileLogger struct {
	*LoggerCore
}

func NewFileLogger(name string, level Level, filepath string) *FileLogger {
	logger := &FileLogger{NewLoggerCore(name, level)}
	logger.AddTransport("defaultFileTransport", NewFileTransport("defaultFileTransport", level, filepath))

	return logger
}
