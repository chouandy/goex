package healthex

// Health health config struct
type Health struct {
	CloudWatchMonitorCfg CloudWatchMonitorCfg `mapstructure:"cloudwatch"`
}
