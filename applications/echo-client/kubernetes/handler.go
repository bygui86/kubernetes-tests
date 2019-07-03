package kubernetes

import (
	"encoding/json"
	"net/http"

	"github.com/bygui86/kubernetes-tests/applications/echo-client/utils/logger"
)

const (
	RESPONSE_STATUS_OK    = "OK"
	RESPONSE_STATUS_ERROR = "ERROR"

	RESPONSE_CODE_OK    = 200
	RESPONSE_CODE_ERROR = 500
)

func livenessHandler(w http.ResponseWriter, r *http.Request) {

	logger.Log.Debugln("[KUBERNETES] Liveness probe invoked...")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		Liveness{
			Status: RESPONSE_STATUS_OK,
			Code:   RESPONSE_CODE_OK,
		},
	)
}

func readynessHandler(w http.ResponseWriter, r *http.Request) {

	logger.Log.Debugln("[KUBERNETES] Readyness probe invoked...")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		Readyness{
			Status: RESPONSE_STATUS_OK,
			Code:   RESPONSE_CODE_OK,
		},
	)
}
