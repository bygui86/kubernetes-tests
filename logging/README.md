# Kubernetes tests logging

## Instructions

### Manual

1. Deploy
	```shell
	kubectl create namespace logging
	kubectl apply -f elasticsearch/elasticsearch_deployment.yaml
	kubectl apply -f elasticsearch/elasticsearch_service.yaml
	kubectl apply -f kibana/kibana_deployment.yaml
	kubectl apply -f kibana/kibana_service.yaml
	kubectl apply -f fluentd/fluentd_rbac.yaml
	kubectl apply -f fluentd/fluentd_daemonset.yaml
	`OPTIONAL` kubectl apply -f fluentd/fluentd_service.yaml
	```

2. Check connection Fluentd-Elastisearch
	You should see that Fluentd connect to Elasticsearch within the logs:
	```shell
	kubectl get pods -n kube-system
	kubectl logs fluentd-* -n kube-system -f
	Connection opened to Elasticsearch cluster => {:host=>"elasticsearch.logging", :port=>9200, :scheme=>"http", :path=>""}
	```

3. Configure Kibana
	1. Open Kibana in browser
		e.g.
		```shell
		minikube service list
			or
		kubectl get svc --all-namespaces
		open http://$(minikube ip):*
		```
	2. Configure using UI
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

4. Back to Kibana
	Left menu: "Discover" to see logs

### Operator

1. Deploy
	1. Add helm charts repository
	```shell
	helm repo add akomljen-charts https://raw.githubusercontent.com/komljen/helm-charts/master/charts/
	```
	2. Install Elasticsearch Operator
	```shell
	helm install --name elasticsearch \
		--namespace logging \
		akomljen-charts/elasticsearch-operator
	```
	3. Install EFK Stack for getting logs from containers and visualise in Kibana
	```shell
	helm install --name efk \
		--namespace logging \
		akomljen-charts/efk
	```

---

## Links

[tutorial-logging](https://mherman.org/blog/logging-in-kubernetes-with-elasticsearch-Kibana-fluentd/)
[fluentd-official](https://docs.fluentd.org/v/0.12/articles/kubernetes-fluentd)
[tutorial-monitoring-loggin](https://medium.com/deepaksood619/ultimate-kubernetes-infrastructure-monitoring-metrics-logs-c7b871d797bd)
