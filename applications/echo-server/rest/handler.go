package rest

import (
	"net/http"

	"github.com/bygui86/kubernetes-tests/applications/echo-server/utils/logger"
	"github.com/gorilla/mux"
)

const (
	DEFAULT_MSG = "Hello world!"
)

func echo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	msg := vars["msg"]

	if len(msg) == 0 {
		logger.Log.Infof("[REST] Echo of default msg '%s'", DEFAULT_MSG)
		w.Write([]byte(DEFAULT_MSG))
	} else {
		logger.Log.Infof("[REST] Echo of msg '%s'", msg)
		w.Write([]byte(msg))
	}
}
