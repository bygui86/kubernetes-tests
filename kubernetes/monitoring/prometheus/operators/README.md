# Kubernetes monitoring - Prometheus - Operators

## Instructions

### Using Helm

1. Deploy Prometheus operator
	```
	helm install --name prometheus \
		--namespace monitoring \
		stable/prometheus-operator
	```

2. Forward port to verify everything properly in place
	* Grafana
		```
		kubectl port-forward svc/prometheus-grafana 3000:80 -n monitoring
		```
	* Prometheus
		```
		kubectl port-forward svc/prometheus-prometheus-oper-prometheus 9090:9090 -n monitoring
		```

### Using plain manifests

1. Deploy Prometheus operator
	```
	kubectl create -f manifests/
	until kubectl get customresourcedefinitions servicemonitors.monitoring.coreos.com ; do date; sleep 1; echo "Waiting for CRDs installation..."; done
	until kubectl get servicemonitors --all-namespaces ; do date; sleep 1; echo "Waiting for ServiceMonitors installation..."; done
	kubectl apply -f manifests/
	```
	`WARN: It can take a few seconds for the above 'create manifests' command to fully create the following resources, so verify the resources are ready before proceeding.`

2. Forward port to verify everything properly in place
	* Grafana
		```
		kubectl port-forward svc/grafana 3000:3000 -n monitoring
		```
	* Prometheus
		```
		kubectl port-forward svc/prometheus-k8s 9090:9090 -n monitoring
		```

---

## Extract Kubernetes manifests

```
cd community-prometheus-operator
helm dependency update
helm template -f values.yaml .
```

---

## Exporters

* [Cassandra](cassandra-exporter)
* [Kafka](kafka-operator)

---

## Links
* [Community-Prometheus-Operator](https://github.com/helm/charts/tree/master/stable/prometheus-operator)
  * [Medium - Tutorial](https://medium.com/deepaksood619/ultimate-kubernetes-infrastructure-monitoring-metrics-logs-c7b871d797bd)
* [CoreOS-Prometheus-Operator](https://github.com/coreos/prometheus-operator)
  * [Getting started](https://github.com/coreos/prometheus-operator/blob/master/Documentation/user-guides/getting-started.md)
  * [How to monitor external services](https://devops.college/prometheus-operator-how-to-monitor-an-external-service-3cb6ac8d5acb)
  * [Alerting](https://github.com/coreos/prometheus-operator/blob/master/Documentation/user-guides/alerting.md)
  * [Alerting examples](https://github.com/coreos/prometheus-operator/tree/master/example/user-guides/alerting)
