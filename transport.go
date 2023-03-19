package lightlog

// ITransport Define an interface called ITransport with several methods
type ITransport interface {
	Enable()
	Disable()
	ShouldLog(level Level) bool
	Log(formattedData string, data *LogData)
	Flush()
	FlushSync()
	Reload()
	Close()
}

type Transport struct {
	name    string // The name of the transport
	level   Level  // The level of the transport
	enabled bool   // The enabled state of the transport
}

// NewTransport Create a new Transport object
func NewTransport(name string, level Level) *Transport {
	return &Transport{
		name:    name,
		level:   level,
		enabled: true,
	}
}

// Enable the transport
func (t *Transport) Enable() {
	t.enabled = true
}

// Disable the transport
func (t *Transport) Disable() {
	t.enabled = false
}

// ShouldLog Check if the transport should log a message with the given level
func (t *Transport) ShouldLog(level Level) bool {
	if !t.enabled {
		return false
	}
	return level >= t.level
}

// Flush log data
func (t *Transport) Flush() {

}

// FlushSync flush log data synchronously
func (t *Transport) FlushSync() {

}

// Reload the transport
func (t *Transport) Reload() {

}

// Close the transport
func (t *Transport) Close() {

}
