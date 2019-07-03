package kubernetes

import (
	"github.com/bygui86/kubernetes-tests/applications/echo-client/utils"
	"github.com/bygui86/kubernetes-tests/applications/echo-client/utils/logger"
)

const (
	// Environment variables -
	kubeHostEnvVar                  = "ECHOCLIENT_KUBE_HOST"
	kubePortEnvVar                  = "ECHOCLIENT_KUBE_PORT"
	kubeServerShutdownTimeoutEnvVar = "ECHOCLIENT_KUBE_SHUTDOWN_TIMEOUT"

	// Default values -
	// host values: '0.0.0.0' for kubernetes, 'localhost' for local
	kubeHostDefault           = "0.0.0.0"
	kubePortDefault           = 7090
	kubeServerShutdownTimeout = 15
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
		ShutdownTimeout: utils.GetIntEnv(kubeServerShutdownTimeoutEnvVar, kubeServerShutdownTimeout),
	}, nil
}
