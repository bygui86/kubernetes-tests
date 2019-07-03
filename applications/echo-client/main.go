package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bygui86/kubernetes-tests/applications/echo-client/kubernetes"
	"github.com/bygui86/kubernetes-tests/applications/echo-client/rest"
	"github.com/bygui86/kubernetes-tests/applications/echo-client/utils/logger"
)

// main -
func main() {

	logger.Log.Infoln("[MAIN] Starting echo-client...")

	kubeServer := startKubernetes()
	defer kubeServer.Shutdown()

	restServer := startRest()
	defer restServer.Shutdown()

	logger.Log.Infoln("[MAIN] echo-client ready!")

	startSysCallChannel()
}

// startKubernetes -
func startKubernetes() *kubernetes.KubeServer {

	server, err := kubernetes.NewKubeServer()
	if err != nil {
		logger.Log.Errorf("[MAIN] Kubernetes server creation failed: %s", err.Error())
		os.Exit(404)
	}
	logger.Log.Debugln("[MAIN] Kubernetes server successfully created")

	server.Start()
	logger.Log.Debugln("[MAIN] Kubernetes successfully started")

	return server
}

// startRest -
func startRest() *rest.RestServer {

	server, err := rest.NewRestServer()
	if err != nil {
		logger.Log.Errorf("[MAIN] REST server creation failed: %s", err.Error())
		os.Exit(404)
	}
	logger.Log.Debugln("[MAIN] REST server successfully created")

	server.Start()
	logger.Log.Debugln("[MAIN] REST successfully started")

	return server
}

// startSysCallChannel -
func startSysCallChannel() {

	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
	logger.Log.Warnln("[MAIN] Termination signal received!")
}
