package rest

import (
	"strconv"

	"github.com/bygui86/kubernetes-tests/applications/client-app-go/utils"
	"github.com/bygui86/kubernetes-tests/applications/client-app-go/utils/logger"
)

const (
	// Environment variables -
	restHostEnvVar                  = "CLIENTGO_REST_HOST"
	restPortEnvVar                  = "CLIENTGO_REST_PORT"
	restServerShutdownTimeoutEnvVar = "CLIENTGO_REST_SHUTDOWN_TIMEOUT"
	restServerAppProtocolEnvVar     = "CLIENTGO_REST_SERVER_APP_PROTOCOL"
	restServerAppHostEnvVar         = "CLIENTGO_REST_SERVER_APP_HOST"
	restServerAppPortEnvVar         = "CLIENTGO_REST_SERVER_APP_PORT"
	restServerAppUrlRootEnvVar      = "CLIENTGO_REST_SERVER_APP_URL_ROOT"

	// Default values -
	// host values: '0.0.0.0' for kubernetes, 'localhost' for local
	restHostDefault                  = "0.0.0.0"
	restPortDefault                  = 8080
	restServerShutdownTimeoutDefault = 15
	restServerAppProtocolDefault     = "http"
	restServerAppHostDefault         = "server-app"
	restServerAppPortDefault         = 8080
	restServerAppUrlRootDefault      = "users"
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
		ShutdownTimeout: utils.GetIntEnv(restServerShutdownTimeoutEnvVar, restServerShutdownTimeoutDefault),
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
