package lightlog

import (
	"fmt"
	"os"
	"time"
)

var pid int
var ip, ipv4, ipv6 string

func init() {
	pid = os.Getpid()
	ipv4, ipv6 = GetIPAddresses()
	ip = ipv4
	if ip == "" {
		ip = ipv6
	}
}

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
	pid       int
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

// scheduleFlush schedules the logger to flush at a given interval
func scheduleFlush(logger *LoggerCore, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		logger.Flush()
	}
}

// NewLoggerCore creates a new LoggerCore instance
func NewLoggerCore(name string, level Level) *LoggerCore {
	logger := &LoggerCore{
		name:       name,
		level:      level,
		transports: make(map[string]ITransport),
	}

	go scheduleFlush(logger, 3*time.Second)

	return logger
}

// AddTransport adds new transport to the logger
func (l *LoggerCore) AddTransport(name string, transport ITransport) {
	if l.GetTransport(name) != nil {
		panic(fmt.Sprintf("transport %s has already exists!", name))
	}
	if transport != nil {
		l.transports[name] = transport
	}
}

// GetTransport returns the transport with the given name
func (l *LoggerCore) GetTransport(name string) ITransport {
	return l.transports[name]
}

// RemoveTransport removes the transport with the given name
func (l *LoggerCore) RemoveTransport(name string) {
	if transport := l.GetTransport(name); transport != nil {
		transport.Close()
		delete(l.transports, name)
	}
}

// DisableTransport disables the transport with the given name
func (l *LoggerCore) DisableTransport(name string) {
	if transport := l.GetTransport(name); transport != nil {
		transport.Disable()
	}
}

// EnableTransport enables the transport with the given name
func (l *LoggerCore) EnableTransport(name string) {
	if transport := l.GetTransport(name); transport != nil {
		transport.Enable()
	}
}

// ReloadTransport reloads the transport with the given name
func (l *LoggerCore) ReloadTransport(name string) {
	if transport := l.GetTransport(name); transport != nil {
		transport.Reload()
	}
}

// ReloadAllTransports reloads all transports
func (l *LoggerCore) ReloadAllTransports() {
	for _, transport := range l.transports {
		transport.Reload()
	}
}

// CloseTransport closes the transport with the given name
func (l *LoggerCore) CloseTransport(name string) {
	if transport := l.GetTransport(name); transport != nil {
		transport.Close()
	}
}

// CloseAllTransport closes all transports
func (l *LoggerCore) CloseAllTransport() {
	for _, transport := range l.transports {
		transport.Close()
	}
}

// Flush flushes all transports
func (l *LoggerCore) Flush() {
	for _, transport := range l.transports {
		transport.Flush()
	}
}

// FlushSync flushes all transports synchronously
func (l *LoggerCore) FlushSync() {
	for _, transport := range l.transports {
		transport.FlushSync()
	}
}

// Close closes all transports
func (l *LoggerCore) Close() {
	l.CloseAllTransport()
}

// Log logs a message with the given level and format
func (l *LoggerCore) Log(level Level, format string, v ...interface{}) {
	if level < l.level {
		return
	}

	levelStr := []string{"TRACE", "DEBUG", "INFO ", "WARN ", "ERROR", "FATAL"}[level]
	message := fmt.Sprintf(format, v...)
	now := time.Now()

	log := &LogData{
		level:     level,
		levelStr:  levelStr,
		time:      now,
		timestamp: now.UnixMilli(),
		datetime:  now.Format("2006-01-02 15:04:05"),
		pid:       pid,
		ip:        ip,
		ipv4:      ipv4,
		ipv6:      ipv6,
		location:  GetLocation(),
		message:   message,
		stack:     "",
		logId:     "",
		tags:      make(map[string]string),
	}
	logMsg := fmt.Sprintf("%s %s %s %s: %s", log.levelStr, log.datetime, log.ip, log.location, log.message)

	for _, transport := range l.transports {
		if transport.ShouldLog(level) {
			transport.Log(logMsg, log)
		}
	}
}

// Trace logs a message with the TRACE level
func (l *LoggerCore) Trace(format string, v ...interface{}) {
	l.Log(TRACE, format, v...)
}

// Debug logs a message with the DEBUG level
func (l *LoggerCore) Debug(format string, v ...interface{}) {
	l.Log(DEBUG, format, v...)
}

// Info logs a message with the INFO level
func (l *LoggerCore) Info(format string, v ...interface{}) {
	l.Log(INFO, format, v...)
}

// Warn logs a message with the WARN level
func (l *LoggerCore) Warn(format string, v ...interface{}) {
	l.Log(WARN, format, v...)
}

// Error logs a message with the ERROR level
func (l *LoggerCore) Error(format string, v ...interface{}) {
	l.Log(ERROR, format, v...)
}

// Fatal logs a message with the FATAL level
func (l *LoggerCore) Fatal(format string, v ...interface{}) {
	l.Log(FATAL, format, v...)
}
