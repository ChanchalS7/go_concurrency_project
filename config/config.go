package config

import "time"

const (
	NumWorkers = 5
	QueueSize = 100
	RateLimit = 200 * time.Millisecond
	MonitorPeriod = 2 *time.Second
	HTTPPort = ":8080"
)