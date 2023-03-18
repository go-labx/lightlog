package lightlog

import (
	"bufio"
	"fmt"
	"os"
)

type FileTransport struct {
	*Transport
	filepath string
	file     *os.File
	writer   *bufio.Writer
}

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

func (f *FileTransport) Log(formattedData string, _ *LogData) {
	_, err := f.writer.WriteString(formattedData)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
}

func (f *FileTransport) Flush() {
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			fmt.Println("Error flushing:", err)
		}
	}(f.writer)
}

func (f *FileTransport) FlushSync() {
	err := f.writer.Flush()
	if err != nil {
		fmt.Println("Error flushing sync:", err)
		return
	}
}

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

func (f *FileTransport) Close() {
	f.FlushSync()

	err := f.file.Close()
	if err != nil {
		fmt.Println("Error closing file:", err)
		return
	}
}
