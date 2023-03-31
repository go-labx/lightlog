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

// Log logs the formatted data to the console
func (c *ConsoleTransport) Log(data *LogData) {
	switch data.level {
	case TRACE:
		fmt.Println(color.CyanString(data.formattedMessage))
	case DEBUG:
		fmt.Println(color.BlueString(data.formattedMessage))
	case INFO:
		fmt.Println(color.GreenString(data.formattedMessage))
	case WARN:
		fmt.Println(color.YellowString(data.formattedMessage))
	case ERROR, FATAL:
		fmt.Println(color.RedString(data.formattedMessage))
	default:
		fmt.Println(data.formattedMessage)
	}
}
