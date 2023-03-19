package lightlog

import "fmt"

type ConsoleTransport struct {
	*Transport
}

func NewConsoleTransport(name string, level Level) *ConsoleTransport {
	return &ConsoleTransport{
		NewTransport(name, level),
	}
}

func (c *ConsoleTransport) Log(formattedData string, _ *LogData) {
	fmt.Println(formattedData)
}
