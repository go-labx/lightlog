# lightlog
a lightweight and high-performance logging library for Go.

## Install

```bash
go get github.com/go-labx/lightlog
```

## Usage

To use `lightlog`, you need to create a logger object. Here's an example of how to create a logger with default options:

```go
options := &LoggerOptions{
    name:     "mylogger",
    level:    INFO,
    filepath: "/var/log/myapp.log",
}
logger := NewLogger(options)
```

Once you have a logger object, you can use it to log messages:

```go
logger.Trace("This is a trace message: %s", message)
logger.Debug("This is a debug message: %s", message)
logger.Info("This is an info message: %s", message)
logger.Warn("This is a warning message: %s", message)
logger.Error("This is an error message: %s", message)
logger.Fatal("This is an fatal message: %s", message)
```

## Transport
`lightlog` provides several transport options for logging messages. By default, `lightlog` creates a console and file transport for each logger. You can also add your own custom transports.

### ConsoleTransport

The ConsoleTransport logs messages to the console. Here's an example of how to create a ConsoleTransport:

```
transport := NewConsoleTransport("myconsole", INFO)
```

### FileTransport

The FileTransport logs messages to a file. Here's an example of how to create a FileTransport:

```
transport := NewFileTransport("myfile", INFO, "/var/log/myapp.log")
```

### Custom Transports

You can also create your own custom transports by implementing the `ITransport` interface. Here's an example of a custom transport:

```
type MyTransport struct {
    *Transport
    // Add any custom fields here
}

func NewMyTransport(name string, level Level) *MyTransport {
    return &MyTransport{
        Transport: NewTransport(name, level),
        // Initialize any custom fields here
    }
}

func (t *MyTransport) Log(formattedData string, data *LogData) {
    // Implement your logging logic here
}
```

## Log Levels

`lightlog` provides several log levels that you can use to filter messages. The available log levels are:

- TRACE
- DEBUG
- INFO
- WARN
- ERROR
- FATAL

To set the log level of a logger, you can use the `level` field of the LoggerOptions struct.