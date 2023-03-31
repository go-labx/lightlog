package lightlog

import (
	"fmt"
	"github.com/go-labx/color"
)

// ConsoleTransport is a struct that represents a console transport
type ConsoleTransport struct {
	*Transport
}

// NewConsoleTransport creates a new ConsoleTransport instance
func NewConsoleTransport(name string, level Level) *ConsoleTransport {
	return &ConsoleTransport{
		NewTransport(name, level),
	}
}

func (c *ConsoleTransport) log(level Level, message string) {
	switch level {
	case TRACE:
		fmt.Println(color.CyanString(message))
	case DEBUG:
		fmt.Println(color.BlueString(message))
	case INFO:
		fmt.Println(color.GreenString(message))
	case WARN:
		fmt.Println(color.YellowString(message))
	case ERROR, FATAL:
		fmt.Println(color.RedString(message))
	default:
		fmt.Println(message)
	}
}

// Log logs the formatted data to the console
func (c *ConsoleTransport) Log(data *LogData) {
	c.log(data.level, data.formattedMessage)
}
