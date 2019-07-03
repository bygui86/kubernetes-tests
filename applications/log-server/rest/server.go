package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/bygui86/kubernetes-tests/applications/log-server/utils/logger"

	"github.com/gorilla/mux"
)

// RestServer -
type RestServer struct {
	Config     *Config
	Router     *mux.Router
	HttpServer *http.Server
}

// NewRestServer - Create new REST server
func NewRestServer() (*RestServer, error) {

	logger.Log.Infoln("[REST] Setup new REST server...")

	// create config
	cfg, err := newConfig()
	if err != nil {
		return nil, err
	}

	// create router
	router := newRouter()

	// create http server
	httpServer := newHttpServer(cfg.RestHost, cfg.RestPort, router)

	return &RestServer{
		Config:     cfg,
		Router:     router,
		HttpServer: httpServer,
	}, nil
}

// newRouter -
func newRouter() *mux.Router {

	logger.Log.Debugln("[REST] Setup new Router config...")

	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/log", log).Methods(http.MethodGet)
	router.HandleFunc("/log/{msg}", log).Methods(http.MethodGet)
	return router
}

// newHttpServer -
func newHttpServer(host string, port int, router *mux.Router) *http.Server {

	logger.Log.Debugf("[REST] Setup new HTTP server on port %d...", port)

	return &http.Server{
		Addr:    host + ":" + strconv.Itoa(port),
		Handler: router,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
}

// Start - Start REST server
func (s *RestServer) Start() {

	logger.Log.Infoln("[REST] Start REST server...")

	// TODO add a channel to communicate if everything is right
	go func() {
		if err := s.HttpServer.ListenAndServe(); err != nil {
			logger.Log.Errorln("[REST] Error starting REST server:", err)
		}
	}()

	logger.Log.Infoln("[REST] REST server listen on port", s.Config.RestPort)
}

// Shutdown - Shutdown REST server
func (s *RestServer) Shutdown() {

	logger.Log.Warnln("[REST] Shutdown REST server...")
	if s.HttpServer != nil {
		// create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.Config.ShutdownTimeout)*time.Second)
		defer cancel()
		// does not block if no connections, otherwise wait until the timeout deadline
		s.HttpServer.Shutdown(ctx)
	}
}
