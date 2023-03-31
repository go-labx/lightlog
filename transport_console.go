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
func (c *ConsoleTransport) Log(formattedData string, data *LogData) {
	switch data.level {
	case TRACE:
		fmt.Println(color.CyanString(formattedData))
	case DEBUG:
		fmt.Println(color.BlueString(formattedData))
	case INFO:
		fmt.Println(color.GreenString(formattedData))
	case WARN:
		fmt.Println(color.YellowString(formattedData))
	case ERROR, FATAL:
		fmt.Println(color.RedString(formattedData))
	default:
		fmt.Println(formattedData)
	}
}
