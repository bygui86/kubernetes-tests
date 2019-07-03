package rest

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/bygui86/kubernetes-tests/applications/echo-client/utils/logger"

	"github.com/gorilla/mux"
)

func (c *ServerAppConfig) echo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	msg := vars["msg"]

	logger.Log.Infof("[REST] Echo of msg '%s'", msg)

	var response *http.Response
	var err error

	if len(msg) == 0 {
		response, err = http.Get(c.Url)
	} else {
		response, err = http.Get(c.Url + "/" + msg)
	}

	if response == nil || err != nil {
		logger.Log.Errorf("[REST] Error 'Echo of msg '%s'': %s", msg, err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer response.Body.Close()
	io.Copy(w, response.Body)
}

func respondWithError(w http.ResponseWriter, code int, message string) {

	w.WriteHeader(code)
	response, _ := json.Marshal(map[string]string{"error": message})
	w.Write(response)
}
