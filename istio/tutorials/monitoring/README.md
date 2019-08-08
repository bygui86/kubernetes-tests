
# # Istio - Monitoring - Prometheus-operator

### Script
```
./istio-monitoring.sh
```

### Manual
1. Deploy Prometheus operator (following solutions are mutually exclusive)
	`WARN: It can take a few seconds for the above 'create manifests' command to fully create the following resources, so verify the resources are ready before proceeding.`
	```
	kubectl apply -f 1_namespace/
	kubectl apply -f 2_crd/
	kubectl apply -f 3_operator/
	kubectl apply -f 4_alert-manager/
	kubectl apply -f 5_grafana/
	kubectl apply -f 6_kube-state-metrics/
	kubectl apply -f 7_node-exporter/
	kubectl apply -f 8_prometheus-adapter/
	kubectl apply -f 9_prometheus/
	kubectl create secret generic prometheus-additional-scrape-configs --from-file=10_istio-monitoring/prometheus-additional.yaml
	kubectl apply -f 10_istio-monitoring/istio-service-monitor.yaml
	kubectl apply -f 10_istio-monitoring/prometheus-updated-clusterRole.yaml
	kubectl apply -f 10_istio-monitoring/prometheus-updated-prometheus.yaml
	kubectl delete pods -l app=prometheus,prometheus=k8s
	kubectl apply -f 10_istio-monitoring/grafana-dashboards/
	kubectl apply -f 10_istio-monitoring/grafana-updated-deployment.yaml
	kubectl delete pods -l app=grafana
	```
2. Forward port to verify everything properly in place
	* Prometheus
		```
		kubectl port-forward svc/prometheus-k8s 9090:9090 -n monitoring
		```
	* Grafana
		```
		kubectl port-forward svc/grafana 3000:3000 -n monitoring
		```

---

## Grafana

### New dashboards
```
kubectl create configmap -n monitoring grafana-dashboard-<JSON_NAME> --from-file=<JSON_PATH>/<JSON_FILENAME>.json --dry-run -o yaml > grafana-dashboard-<JSON_NAME>.yaml
```

### Istio dashboards
```
kubectl create configmap -n monitoring grafana-dashboard-galley-dashboard --from-file=10_istio-monitoring/grafana-dashboards/json/galley-dashboard.json --dry-run -o yaml > grafana-dashboard-galley-dashboard.yaml
kubectl create configmap -n monitoring grafana-dashboard-istio-mesh-dashboard --from-file=10_istio-monitoring/grafana-dashboards/json/istio-mesh-dashboard.json --dry-run -o yaml > grafana-dashboard-istio-mesh-dashboard.yaml
kubectl create configmap -n monitoring grafana-dashboard-istio-performance-dashboard --from-file=10_istio-monitoring/grafana-dashboards/json/istio-performance-dashboard.json --dry-run -o yaml > grafana-dashboard-istio-performance-dashboard.yaml
kubectl create configmap -n monitoring grafana-dashboard-istio-service-dashboard --from-file=10_istio-monitoring/grafana-dashboards/json/istio-service-dashboard.json --dry-run -o yaml > grafana-dashboard-istio-service-dashboard.yaml
kubectl create configmap -n monitoring grafana-dashboard-istio-workload-dashboard --from-file=10_istio-monitoring/grafana-dashboards/json/istio-workload-dashboard.json --dry-run -o yaml > grafana-dashboard-istio-workload-dashboard.yaml
kubectl create configmap -n monitoring grafana-dashboard-mixer-dashboard --from-file=10_istio-monitoring/grafana-dashboards/json/mixer-dashboard.json --dry-run -o yaml > grafana-dashboard-mixer-dashboard.yaml
kubectl create configmap -n monitoring grafana-dashboard-pilot-dashboard --from-file=10_istio-monitoring/grafana-dashboards/json/pilot-dashboard.json --dry-run -o yaml > grafana-dashboard-pilot-dashboard.yaml
```
