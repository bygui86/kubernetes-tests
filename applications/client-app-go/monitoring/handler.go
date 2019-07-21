package monitoring

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func getMetricsHandler() http.Handler {

	return promhttp.Handler()
}
