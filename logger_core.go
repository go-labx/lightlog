package lightlog

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

type Level int

const (
	TRACE Level = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

type LoggerCore struct {
	name       string
	level      Level
	transports map[string]ITransport
}

type LogData struct {
	level     Level
	levelStr  string
	time      time.Time
	timestamp int64
	pid       int32
	datetime  string
	ipv4      string
	ipv6      string
	ip        string
	location  string
	message   string
	stack     string
	logId     string
	tags      map[string]string
}

func NewLoggerCore(name string, level Level) *LoggerCore {
	return &LoggerCore{
		name:       name,
		level:      level,
		transports: make(map[string]ITransport, 0),
	}
}

func (l *LoggerCore) AddTransport(name string, transport ITransport) {
	if l.GetTransport(name) != nil {
		panic(fmt.Sprintf("transport %s already exists!", name))
	}
	l.transports[name] = transport
}

func (l *LoggerCore) GetTransport(name string) ITransport {
	return l.transports[name]
}

func (l *LoggerCore) RemoveTransport(name string) {
	transport := l.transports[name]
	if transport != nil {
		transport.Close()
		delete(l.transports, name)
	}
}

func (l *LoggerCore) DisableTransport(name string) {
	transport := l.GetTransport(name)
	if transport != nil {
		transport.Disable()
	}
}

func (l *LoggerCore) EnableTransport(name string) {
	transport := l.GetTransport(name)
	if transport != nil {
		transport.Enable()
	}
}

func (l *LoggerCore) ReloadTransport(name string) {
	transport := l.GetTransport(name)
	if transport != nil {
		transport.Reload()
	}
}

func (l *LoggerCore) ReloadAllTransports() {
	for _, transport := range l.transports {
		transport.Reload()
	}
}

func (l *LoggerCore) CloseTransport(name string) {
	transport := l.GetTransport(name)
	if transport != nil {
		transport.Close()
	}
}

func (l *LoggerCore) CloseAllTransport() {
	for _, transport := range l.transports {
		transport.Close()
	}
}

func (l *LoggerCore) Flush() {
	for _, transport := range l.transports {
		transport.Flush()
	}
}

func (l *LoggerCore) FlushSync() {
	for _, transport := range l.transports {
		transport.FlushSync()
	}
}

func (l *LoggerCore) Close() {
	l.FlushSync()
	l.CloseAllTransport()
}

func (l *LoggerCore) Log(level Level, format string, v ...interface{}) {
	if level < l.level {
		return
	}

	levelStr := []string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}[level]
	message := fmt.Sprintf(format, v...)
	now := time.Now()

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	} else {
		file = filepath.Base(file)
	}
	location := fmt.Sprintf("%s:%d", file, line)
	ipv4, ipv6 := GetIPAddresses()
	ip := ipv4
	if ip == "" {
		ip = ipv6
	}

	log := &LogData{
		level:     level,
		levelStr:  levelStr,
		time:      now,
		timestamp: now.UnixMilli(),
		datetime:  now.Format("2006-01-02 15:04:05"),
		pid:       0,
		ip:        ip,
		ipv4:      ipv4,
		ipv6:      ipv6,
		location:  location,
		message:   message,
		stack:     "",
		logId:     "",
		tags:      make(map[string]string),
	}
	logMsg := fmt.Sprintf("%s %s %s %s %s", log.levelStr, log.datetime, log.location, log.ip, log.message)

	for _, transport := range l.transports {
		transport.Log(logMsg, log)
	}
}

func (l *LoggerCore) Trace(format string, v ...interface{}) {
	l.Log(TRACE, format, v...)
}

func (l *LoggerCore) Debug(format string, v ...interface{}) {
	l.Log(DEBUG, format, v...)
}

func (l *LoggerCore) Info(format string, v ...interface{}) {
	l.Log(INFO, format, v...)
}

func (l *LoggerCore) Warn(format string, v ...interface{}) {
	l.Log(WARN, format, v...)
}

func (l *LoggerCore) Error(format string, v ...interface{}) {
	l.Log(ERROR, format, v...)
}

func (l *LoggerCore) Fatal(format string, v ...interface{}) {
	l.Log(FATAL, format, v...)
}
