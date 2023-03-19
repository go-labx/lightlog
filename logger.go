package lightlog

type Logger struct {
	*LoggerCore
}

type LoggerOptions struct {
	Name     string
	Level    Level
	Filepath string
}

func NewLogger(options *LoggerOptions) *Logger {
	if options.Name == "" {
		options.Name = "defaultLogger"
	}

	logger := &Logger{
		NewLoggerCore(options.Name, options.Level),
	}
	logger.AddTransport("defaultConsoleTransport", NewConsoleTransport("defaultConsoleTransport", options.Level))

	if options.Filepath == "" {
		logger.Warn("`options.Filepath` cannot be empty")
	} else {
		logger.AddTransport("defaultFileTransport", NewFileTransport("defaultFileTransport", options.Level, options.Filepath))
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
