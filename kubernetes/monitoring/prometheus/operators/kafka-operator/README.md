# Kubernetes monitoring - Prometheus - Kafka operator

## Instructions

### Download

* Download Strimzi Kafka operator
	```
	curl -L https://github.com/strimzi/strimzi-kafka-operator/releases/download/0.12.1/strimzi-cluster-operator-0.12.1.yaml | sed 's/namespace: .*/namespace: kafka/' > strimzi-kafka-operator.yaml
	```

* Download Kafka cluster
	```
	curl -L https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/0.12.1/examples/kafka/kafka-persistent-single.yaml > strimzi-kafka-cluster.yaml
	```

* Download Kafka cluster configs
	```
	curl -L https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/release-0.12.x/metrics/examples/prometheus/additional-properties/prometheus-additional.yaml > additional-metrics.yaml
	curl -L https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/release-0.12.x/metrics/examples/prometheus/install/strimzi-service-monitor.yaml > strimzi-kafka-service-monitor.yaml
	curl -L https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/release-0.12.x/metrics/examples/prometheus/install/prometheus-rules.yaml | sed -e "s/namespace: .*/namespace: kafka/" > strimzi-kafka-alerting-rules.yaml
	```

### Deployment

1. Deploy Strimzi Kafka operator
	```
	kubectl apply -f 1_strimzi-kafka-operator.yaml
	```

2. Deploy Kafka cluster
	* minikube
		```
		kubectl apply -f 2-mk_strimzi-kafka-cluster.yaml
		```
	* gke
		```
		kubectl apply -f 2-gke_storageclass.yaml
		kubectl apply -f 2-gke_strimzi-kafka-cluster.yaml
		```

3. Deploy Prometheus additional metrics as secret
	```
	kubectl create secret generic additional-scrape-configs --from-file=3_additional-metrics.yaml -n <KAFKA_CLUSTER_NAMESPACE>
	kubectl create secret generic additional-scrape-configs --from-file=3_additional-metrics.yaml -n <PROMETHEUS_NAMESPACE>
	```

4. Deploy Prometheus ServiceMonitor
	```
	kubectl apply -f 4_strimzi-kafka-service-monitor.yaml
	```

5. Deploy Alerting PrometheusRules
	```
	kubectl apply -f 5_strimzi-kafka-alerting-rules.yaml
	```

---

## Links
* [Strimzi Kafka Operator](https://github.com/strimzi/strimzi-kafka-operator/)
  * [Minikube quick start](https://strimzi.io/quickstarts/minikube/)
