package rest

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/bygui86/kubernetes-tests/applications/client-app-go/utils/logger"

	"github.com/gorilla/mux"
)

func (c *ServerAppConfig) getAll(w http.ResponseWriter, r *http.Request) {

	logger.Log.Infoln("[REST] Get all users")

	w.Header().Set("Content-Type", "application/json")
	response, err := http.Get(c.Url)
	if response == nil || err != nil {
		logger.Log.Errorln("[REST] Error 'Get all users':", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	io.Copy(w, response.Body)
	response.Body.Close()
}

func (c *ServerAppConfig) getByEmail(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	email := vars["email"]
	if len(email) == 0 {
		logger.Log.Errorln("[REST] Error 'get user by email': invalid email")
		respondWithError(w, http.StatusInternalServerError, "Invalid email")
		return
	}

	logger.Log.Infoln("[REST] Get user by email", email)

	w.Header().Set("Content-Type", "application/json")
	response, err := http.Get(c.Url + "/" + email)
	if response == nil || err != nil {
		logger.Log.Errorf("[REST] Error 'Get user by email %s': %s", email, err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer response.Body.Close()
	io.Copy(w, response.Body)
}

func (c *ServerAppConfig) insert(w http.ResponseWriter, r *http.Request) {

	logger.Log.Infoln("[REST] Insert new user")

	w.Header().Set("Content-Type", "application/json")
	response, err := http.Post(c.Url, "application/json", r.Body)
	if response == nil || err != nil {
		logger.Log.Errorln("[REST] Error 'Get all users':", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer response.Body.Close()
	io.Copy(w, response.Body)
}

func (c *ServerAppConfig) update(w http.ResponseWriter, r *http.Request) {

	logger.Log.Infoln("[REST] Update user")

	request, requestErr := http.NewRequest(http.MethodPut, c.Url, r.Body)
	if requestErr != nil {
		logger.Log.Errorln("[REST] Error 'Update user':", requestErr.Error())
		respondWithError(w, http.StatusInternalServerError, requestErr.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	response, responseErr := httpExecRequest(request)
	if responseErr != nil {
		logger.Log.Errorln("[REST] Error 'Update user':", responseErr.Error())
		respondWithError(w, http.StatusInternalServerError, responseErr.Error())
		return
	}
	defer response.Body.Close()
	io.Copy(w, response.Body)
}

func (c *ServerAppConfig) deleteAll(w http.ResponseWriter, r *http.Request) {

	logger.Log.Infoln("[REST] Delete all users")

	request, requestErr := http.NewRequest(http.MethodDelete, c.Url, nil)
	if requestErr != nil {
		logger.Log.Errorln("[REST] Error 'Delete all users':", requestErr.Error())
		respondWithError(w, http.StatusInternalServerError, requestErr.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response, responseErr := httpExecRequest(request)
	if responseErr != nil {
		logger.Log.Errorln("[REST] Error 'Delete all users':", responseErr.Error())
		respondWithError(w, http.StatusInternalServerError, responseErr.Error())
		return
	}
	defer response.Body.Close()
	io.Copy(w, response.Body)
}

func (c *ServerAppConfig) deleteByEmail(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	email := vars["email"]
	if len(email) == 0 {
		logger.Log.Errorln("[REST] Error 'Delete user by email': invalid email")
		respondWithError(w, http.StatusInternalServerError, "Invalid email")
		return
	}

	logger.Log.Infoln("[REST] Delete user by email", email)

	request, requestErr := http.NewRequest(http.MethodDelete, c.Url+"/"+email, nil)
	if requestErr != nil {
		logger.Log.Errorf("[REST] Error 'Delete user by email %s': %s", email, requestErr.Error())
		respondWithError(w, http.StatusInternalServerError, requestErr.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response, responseErr := httpExecRequest(request)
	if responseErr != nil {
		logger.Log.Errorf("[REST] Error 'Delete user by email %s': %s", email, responseErr.Error())
		respondWithError(w, http.StatusInternalServerError, responseErr.Error())
		return
	}
	defer response.Body.Close()
	io.Copy(w, response.Body)
}

func httpExecRequest(request *http.Request) (*http.Response, error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func respondWithError(w http.ResponseWriter, code int, message string) {

	w.WriteHeader(code)
	response, _ := json.Marshal(map[string]string{"error": message})
	w.Write(response)
}
