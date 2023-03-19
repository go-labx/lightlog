package main

import (
	"github.com/go-labx/lightlog"
	"net/http"
)

func main() {
	logger := lightlog.NewLogger(&lightlog.LoggerOptions{
		Name:     "logger",
		Level:    lightlog.INFO,
		Filepath: "./logs/1.log",
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		logger.Info("new request --->")
		_, err := writer.Write([]byte("Hello World!"))
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
