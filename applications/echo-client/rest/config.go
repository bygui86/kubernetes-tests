package rest

import (
	"strconv"

	"github.com/bygui86/kubernetes-tests/applications/echo-client/utils"
	"github.com/bygui86/kubernetes-tests/applications/echo-client/utils/logger"
)

const (
	// Environment variables -
	restHostEnvVar              = "ECHOCLIENT_REST_HOST"
	restPortEnvVar              = "ECHOCLIENT_REST_PORT"
	restShutdownTimeoutEnvVar   = "ECHOCLIENT_REST_SHUTDOWN_TIMEOUT"
	restServerAppProtocolEnvVar = "ECHOCLIENT_REST_SERVER_APP_PROTOCOL"
	restServerAppHostEnvVar     = "ECHOCLIENT_REST_SERVER_APP_HOST"
	restServerAppPortEnvVar     = "ECHOCLIENT_REST_SERVER_APP_PORT"
	restServerAppUrlRootEnvVar  = "ECHOCLIENT_REST_SERVER_APP_URL_ROOT"

	// Default values -
	// host values: '0.0.0.0' for kubernetes, 'localhost' for local
	restHostDefault              = "0.0.0.0"
	restPortDefault              = 7001
	restShutdownTimeoutDefault   = 15
	restServerAppProtocolDefault = "http"
	restServerAppHostDefault     = "echo-server"
	restServerAppPortDefault     = 7001
	restServerAppUrlRootDefault  = "echo"
)

// Config -
type Config struct {
	RestHost        string
	RestPort        int
	ShutdownTimeout int
	ServerAppConfig *ServerAppConfig
}

// ServerAppConfig -
type ServerAppConfig struct {
	Protocol string
	Host     string
	Port     int
	UrlRoot  string
	Url      string
}

// newConfig -
func newConfig() (*Config, error) {

	logger.Log.Debugln("[REST] Setup new REST server config...")

	protocol := utils.GetStringEnv(restServerAppProtocolEnvVar, restServerAppProtocolDefault)
	host := utils.GetStringEnv(restServerAppHostEnvVar, restServerAppHostDefault)
	port := utils.GetIntEnv(restServerAppPortEnvVar, restServerAppPortDefault)
	urlRoot := utils.GetStringEnv(restServerAppUrlRootEnvVar, restServerAppUrlRootDefault)
	url := buildServerAppUrl(protocol, host, urlRoot, port)
	return &Config{
		RestHost:        utils.GetStringEnv(restHostEnvVar, restHostDefault),
		RestPort:        utils.GetIntEnv(restPortEnvVar, restPortDefault),
		ShutdownTimeout: utils.GetIntEnv(restShutdownTimeoutEnvVar, restShutdownTimeoutDefault),
		ServerAppConfig: &ServerAppConfig{
			Protocol: protocol,
			Host:     host,
			Port:     port,
			UrlRoot:  urlRoot,
			Url:      url,
		},
	}, nil
}

// buildServerAppUrl -
func buildServerAppUrl(protocol, host, urlRoot string, port int) string {

	return protocol + "://" + host + ":" + strconv.Itoa(port) + "/" + urlRoot
}
