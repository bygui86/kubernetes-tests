# # Kubernetes monitoring - Prometheus - Cassandra exporter

## Instructions

1. Deploy Prometheus operator

2. Deploy Cassandra together with Cassandra Exporter for Prometheus
	```
	kubectl apply -f 1_cassandra-configmap.yaml
	kubectl apply -f 2_exporter-configmap.yaml
	kubectl apply -f 3_statefulset.yaml
	kubectl apply -f 4_service.yaml
	```

3. Setup Prometheus to discover Cassandra Exporter (through ServiceMonitor)
	```
	kubectl apply -f 5_exporter-servicemonitor.yaml
	```

4. Setup Prometheus alerts for Cassandra (through PrometheusRules)
	```
	kubectl apply -f 6_exporter-alerting-rules.yaml
	```

`WARN`
Use following command to retrieve the right labels for the ServiceMonitor 
`k get prometheus mon-prometheus-operator-prometheus -o yaml | jq .spec.serviceMonitorSelector`
and the PrometheusRules
`k get prometheus mon-prometheus-operator-prometheus -o yaml | jq .spec.ruleSelector`

---

## Links
* [Cassandra-Exporter](https://github.com/criteo/cassandra_exporter)
* [MySocialApp-Cassandra-Helm-Chart](https://github.com/MySocialApp/kubernetes-helm-chart-cassandra)
