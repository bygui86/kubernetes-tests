# Kubernetes monitoring - Prometheus

## Instructions

### Operator

1. Deploy
	```shell
	helm install --name prometheus \
		--namespace monitoring \
		stable/prometheus-operator
	```

2. Forward port
	* Grafana
	```shell
		kubectl port-forward svc/prometheus-grafana 3000:80 -n monitoring
	```
	* Prometheus
	```shell
	kubectl port-forward svc/prometheus-prometheus-oper-prometheus 9090:9090 -n monitoring
	```

### Manual

`TODO`

---

## Links
* [prometheus-opeartor](https://medium.com/deepaksood619/ultimate-kubernetes-infrastructure-monitoring-metrics-logs-c7b871d797bd)
* [manual-install](https://vadosware.io/post/better-k8s-monitoring-part-1-adding-prometheus/)