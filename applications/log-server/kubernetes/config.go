package kubernetes

import (
	"github.com/bygui86/kubernetes-tests/applications/log-server/utils"
	"github.com/bygui86/kubernetes-tests/applications/log-server/utils/logger"
)

const (
	// Environment variables -
	kubeHostEnvVar            = "LOGSERVER_KUBE_HOST"
	kubePortEnvVar            = "LOGSERVER_KUBE_PORT"
	kubeShutdownTimeoutEnvVar = "LOGSERVER_KUBE_SHUTDOWN_TIMEOUT"

	// Default values -
	// host values: '0.0.0.0' for kubernetes, 'localhost' for local
	kubeHostDefault     = "0.0.0.0"
	kubePortDefault     = 7090
	kubeShutdownTimeout = 15
)

// Config -
type Config struct {
	RestHost        string
	RestPort        int
	ShutdownTimeout int
}

// newConfig -
func newConfig() (*Config, error) {

	logger.Log.Debugln("[KUBERNETES] Setup new Kubernetes config...")

	return &Config{
		RestHost:        utils.GetStringEnv(kubeHostEnvVar, kubeHostDefault),
		RestPort:        utils.GetIntEnv(kubePortEnvVar, kubePortDefault),
		ShutdownTimeout: utils.GetIntEnv(kubeShutdownTimeoutEnvVar, kubeShutdownTimeout),
	}, nil
}
