
# Istio bookinfo - Telemetry

## Metrics

1. Apply new metrics manifest. Istio will generate and collect them automatically.
	```
	kubectl apply -f samples/bookinfo/telemetry/metrics.yaml
	```

2. Generate some traffic making requests to `http://$GATEWAY_URL/productpage`

3. Verify that the new metric values are being generated and collected.
	1. Open a new cli tab
		```
		kubectl port-forward svc/prometheus 9090:9090 -n istio-system
		```
	2. Open a new browser tab
		```
		http://localhost:9090/graph?g0.range_input=1h&g0.expr=istio_double_request_count&g0.tab=1
		```

4. Visualizing metrics
	1. Open a new cli tab
		```
		kubectl port-forward svc/grafana 3000:3000 -n istio-system
		```
	2. Open some other browser tabs
		```
		http://localhost:3000/dashboard/db/istio-mesh-dashboard
		http://localhost:3000/dashboard/db/istio-service-dashboard
		http://localhost:3000/dashboard/db/istio-workload-dashboard
		```

5. Generate some other traffic making requests to `http://$GATEWAY_URL/productpage` to better visualize graphs

## Logs

`TODO`

## Tracing

`TODO`

---

## Links
* [Telemetry](https://istio.io/docs/tasks/telemetry/)
  * Metrics
    * [Collecting metrics](https://istio.io/docs/tasks/telemetry/metrics/collecting-metrics/)
    * `TODO` [Collecting TCP metrics](https://istio.io/docs/tasks/telemetry/metrics/tcp-metrics/)
    * [Querying metrics](https://istio.io/docs/tasks/telemetry/metrics/querying-metrics/)
    * [Visualizing metrics](https://istio.io/docs/tasks/telemetry/metrics/using-istio-dashboard/)
  * `TODO` Logs
    * `TODO` [Collecting logs](https://istio.io/docs/tasks/telemetry/logs/collecting-logs/)
    * `TODO` [Getting Envoy's access logs](https://istio.io/docs/tasks/telemetry/logs/access-log/)
    * `TODO` [Logging with Fluentd](https://istio.io/docs/tasks/telemetry/logs/fluentd/)
  * Distributed tracing
    * `TODO` [Overview](https://istio.io/docs/tasks/telemetry/distributed-tracing/overview/)
    * `TODO` [Jaeger](https://istio.io/docs/tasks/telemetry/distributed-tracing/jaeger/)
   * `TODO` [Visualising the mesh (Kiali)](https://istio.io/docs/tasks/telemetry/kiali/)
   * `TODO` [Remote accessing telemetry addon](https://istio.io/docs/tasks/telemetry/gateways/)
