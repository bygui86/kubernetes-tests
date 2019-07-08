# Kubernetes tests - Applications - echo-server

## description
Simple application written in Golang that echoes all messages received tughour exposed via REST APIs

## ports
* 7001		rest apis
* 7090		rest apis for kubernetes probes

## rest apis
* port 7001
  * GET /echo			echo default message "Hello Wolrd"
  * GET /echo/{msg}		echo back incoming message
* port 7090
  * GET /live		kubernetes liveness probe
  * GET /ready		kubernetes readyness probe
