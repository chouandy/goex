package healthex

// Health health config struct
type Health struct {
	CloudWatchMonitor CloudWatchMonitor `mapstructure:"cloudwatch_monitor"`
}
