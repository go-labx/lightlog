package lightlog

// Logger is a struct that contains a pointer to LoggerCore
type Logger struct {
	*LoggerCore
}

// LoggerOptions is a struct that contains options for creating a new Logger
type LoggerOptions struct {
	Name     string // Name of the logger
	Level    Level // Level of the logger
	Filepath string // Filepath for the logger
}

// NewLogger creates a new Logger with the given options
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

// ConsoleLogger is a struct that contains a pointer to LoggerCore
type ConsoleLogger struct {
	*LoggerCore
}

// NewConsoleLogger creates a new ConsoleLogger with the given name and level
func NewConsoleLogger(name string, level Level) *ConsoleLogger {
	logger := &ConsoleLogger{NewLoggerCore(name, level)}
	logger.AddTransport("defaultConsoleTransport", NewConsoleTransport("defaultConsoleTransport", level))

	return logger
}

// FileLogger is a struct that contains a pointer to LoggerCore
type FileLogger struct {
	*LoggerCore
}

// NewFileLogger creates a new FileLogger with the given name, level, and filepath
func NewFileLogger(name string, level Level, filepath string) *FileLogger {
	logger := &FileLogger{NewLoggerCore(name, level)}
	logger.AddTransport("defaultFileTransport", NewFileTransport("defaultFileTransport", level, filepath))

	return logger
}
