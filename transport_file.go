package lightlog

import (
	"bufio"
	"fmt"
	"os"
)

// FileTransport is a struct that represents a transport for logging to a file.
type FileTransport struct {
	*Transport
	filepath string
	file     *os.File
	writer   *bufio.Writer
}

// NewFileTransport is a constructor function that creates a new instance of FileTransport.
func NewFileTransport(name string, level Level, filepath string) *FileTransport {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}

	writer := bufio.NewWriter(file)

	return &FileTransport{
		NewTransport(name, level),
		filepath,
		file,
		writer,
	}
}

// Log is a method that writes formattedData to the file.
func (f *FileTransport) Log(formattedData string, _ *LogData) {
	_, err := f.writer.WriteString(formattedData)
	_, err = f.writer.WriteString("\n")

	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
}

// Flush is a method that flushes the writer's buffer to the file.
func (f *FileTransport) Flush() {
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			fmt.Println("Error flushing:", err)
		}
	}(f.writer)
}

// FlushSync is a method that synchronously flushes the writer's buffer to the file.
func (f *FileTransport) FlushSync() {
	err := f.writer.Flush()
	if err != nil {
		fmt.Println("Error flushing sync:", err)
		return
	}
}

// Reload is a method that flushes the writer's buffer to the file, reloads the file, and closes the old file.
func (f *FileTransport) Reload() {
	f.FlushSync()

	file, err := os.OpenFile(f.filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error reloading file:", err)
		return
	}
	f.Close()

	f.file = file
}

// Close is a method that flushes the writer's buffer to the file and closes the file.
func (f *FileTransport) Close() {
	f.FlushSync()

	err := f.file.Close()
	if err != nil {
		fmt.Println("Error closing file:", err)
		return
	}
}
