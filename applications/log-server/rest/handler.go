package rest

import (
	"net/http"

	"github.com/bygui86/kubernetes-tests/applications/log-server/utils/logger"
	"github.com/gorilla/mux"

	"gopkg.in/matryer/respond.v1"
)

const (
	DEFAULT_MSG = "Hello world!"
)

func log(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	msg := vars["msg"]

	logger.Log.Infof("[REST] Message received: %s", msg)
	respond.With(w, r, http.StatusOK, nil)
}
