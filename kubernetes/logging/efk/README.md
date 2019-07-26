
# Kubernetes tests - Logging - Elastisearch, Fluentd and Kibana

## Instructions

1. Deploy
	```
	minikube start --cpus 4 --memory 8192

	kubectl create namespace logging

	kubectl apply -f kubernetes/elastic.yaml -n logging
	kubectl apply -f kubernetes/kibana.yaml -n logging
	kubectl apply -f kubernetes/fluentd-rbac.yaml
	`OPTIONAL` kubectl apply -f kubernetes/fluentd-service.yaml
	kubectl apply -f kubernetes/fluentd-daemonset.yaml

	kubectl get pods -n kube-system
	kubectl logs fluentd-* -n kube-system -f
	```

2. Check connection Fluentd-Elastisearch
	You should see that Fluentd connect to Elasticsearch within the logs:
	```
	Connection opened to Elasticsearch cluster => {:host=>"elasticsearch.logging", :port=>9200, :scheme=>"http", :path=>""}
	```

1. Configure Kibana
	```
	minikube service list
		or
	kubectl get svc --all-namespaces
	open http://$(minikube ip):*
	```
	1. Left menu: "Management"
	2. Kibana group: "Index Patterns"
	3. Button: "Create index pattern"
	4. Index creation form 1:
		Index pattern: logstash*
	5. Button: "Next step"
	6. Index creation form 2:
		Time Filter field name: @timestamp
	7. Button: "Create index pattern"
	8. Leave open the window

4. Spin up application
	```
	docker build -t fluentd-node-sample:latest -f sample-app/Dockerfile sample-app
	kubectl apply -f kubernetes/node-deployment.yaml -n default
	```

5. Back to Kibana
	Left menu: "Discover" to see logs

---

## Links
* [logging-tutorial-1](https://mherman.org/blog/logging-in-kubernetes-with-elasticsearch-Kibana-fluentd/)
* [logging-tutorial-2](https://vadosware.io/post/better-k8s-monitoring-part-2-adding-logging-with-efkk/)
* [fluentd-official](https://docs.fluentd.org/v/0.12/articles/kubernetes-fluentd)
* [monitoring-logging-tutorial](https://medium.com/deepaksood619/ultimate-kubernetes-infrastructure-monitoring-metrics-logs-c7b871d797bd)
