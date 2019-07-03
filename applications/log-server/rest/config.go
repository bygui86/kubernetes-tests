package rest

import (
	"github.com/bygui86/kubernetes-tests/applications/log-server/utils"
	"github.com/bygui86/kubernetes-tests/applications/log-server/utils/logger"
)

const (
	// Environment variables -
	restHostEnvVar            = "LOGSERVER_REST_HOST"
	restPortEnvVar            = "LOGSERVER_REST_PORT"
	restShutdownTimeoutEnvVar = "LOGSERVER_REST_SHUTDOWN_TIMEOUT"

	// Default values -
	// host values: '0.0.0.0' for kubernetes, 'localhost' for local
	restHostDefault            = "0.0.0.0"
	restPortDefault            = 7001
	restShutdownTimeoutDefault = 15
)

// Config -
type Config struct {
	RestHost        string
	RestPort        int
	ShutdownTimeout int
}

// newConfig -
func newConfig() (*Config, error) {

	logger.Log.Debugln("[REST] Setup new REST server config...")

	return &Config{
		RestHost:        utils.GetStringEnv(restHostEnvVar, restHostDefault),
		RestPort:        utils.GetIntEnv(restPortEnvVar, restPortDefault),
		ShutdownTimeout: utils.GetIntEnv(restShutdownTimeoutEnvVar, restShutdownTimeoutDefault),
	}, nil
}
