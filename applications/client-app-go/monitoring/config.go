package monitoring

import (
	"github.com/bygui86/kubernetes-tests/applications/client-app-go/utils"
	"github.com/bygui86/kubernetes-tests/applications/client-app-go/utils/logger"
)

const (
	// Environment variables -
	monHostEnvVar            = "CLIENTGO_MONITOR_HOST"
	monPortEnvVar            = "CLIENTGO_MONITOR_PORT"
	monShutdownTimeoutEnvVar = "CLIENTGO_MONITOR_SHUTDOWN_TIMEOUT"

	// Default values -
	// host values: '0.0.0.0' for kubernetes, 'localhost' for local
	monHostDefault     = "localhost"
	monPortDefault     = 8091
	monShutdownTimeout = 15
)

// Config -
type Config struct {
	RestHost        string
	RestPort        int
	ShutdownTimeout int
}

// newConfig -
func newConfig() (*Config, error) {

	logger.Log.Debugln("[MONITORING] Setup new Monitoring config...")

	return &Config{
		RestHost:        utils.GetStringEnv(monHostEnvVar, monHostDefault),
		RestPort:        utils.GetIntEnv(monPortEnvVar, monPortDefault),
		ShutdownTimeout: utils.GetIntEnv(monShutdownTimeoutEnvVar, monShutdownTimeout),
	}, nil
}
