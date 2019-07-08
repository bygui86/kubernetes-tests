# Kubernetes tests - Applications - log-server

## description
Simple application written in Golang that logs all messages received tughour exposed via REST APIs

## ports
* 7001		rest apis
* 7090		rest apis for kubernetes probes

## rest apis
* port 7001
  * GET /log		log default message "Hello Wolrd"
  * GET /log/{msg}	log incoming message
* port 7090
  * GET /live		kubernetes liveness probe
  * GET /ready		kubernetes readyness probe
