package lightlog

type ITransport interface {
	Enable()
	Disable()
	Log(formattedData string, data *LogData)
	Flush()
	FlushSync()
	Reload()
	Close()
}

type Transport struct {
	name    string
	level   Level
	enabled bool
}

func NewTransport(name string, level Level) *Transport {
	return &Transport{
		name:    name,
		level:   level,
		enabled: true,
	}
}

func (t *Transport) Enable() {
	t.enabled = true
}

func (t *Transport) Disable() {
	t.enabled = false
}

func (t *Transport) ShouldLog(level Level) bool {
	if !t.enabled {
		return false
	}
	return level > t.level
}

func (t *Transport) Flush() {

}

func (t *Transport) FlushSync() {

}

func (t *Transport) Reload() {

}

func (t *Transport) Close() {

}
