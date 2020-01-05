
# Kubernetes tests - Logging - Elastisearch, Fluentd and Kibana

## Pre-requisites

* cpu: 4
* memory: 8192

---

## Instructions

1. Deploy
	```
	kubectl apply -f .
	```

2. Check connection Fluentd-Elastisearch
	```
	kubectl get pods -n kube-system
	kubectl logs fluentd-* -n kube-system -f
	```
	You should see that Fluentd connect to Elasticsearch within the logs:
	```
	Connection opened to Elasticsearch cluster => {:host=>"elasticsearch.logging", :port=>9200, :scheme=>"http", :path=>""}
	```

3. Configure Kibana
	```
	open http://$(minikube ip):$(kubectl get svc -n logging | grep -i kibana | awk '{print $5}' | cut -d ',' -f 1 | sed 's,[0-9]*:,,' | sed 's,/TCP,,')
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
* docs
  * [fluentd-official](https://docs.fluentd.org/v/0.12/articles/kubernetes-fluentd)
* tutorials
  * [logging-tutorial-1](https://mherman.org/blog/logging-in-kubernetes-with-elasticsearch-Kibana-fluentd/)
  * [logging-tutorial-2](https://vadosware.io/post/better-k8s-monitoring-part-2-adding-logging-with-efkk/)
  * [monitoring-logging-tutorial](https://medium.com/deepaksood619/ultimate-kubernetes-infrastructure-monitoring-metrics-logs-c7b871d797bd)
  * [digital-ocean-tutorial](https://www.digitalocean.com/community/tutorials/how-to-set-up-an-elasticsearch-fluentd-and-kibana-efk-logging-stack-on-kubernetes#step-2-%E2%80%94-creating-the-elasticsearch-statefulset)
* operator
  * [official](https://www.elastic.co/products/elastic-cloud-kubernetes)
* samples
  * [neogenix](https://github.com/neogenix/k8s-elk)
  * [giantswarm](https://github.com/giantswarm/kubernetes-elastic-stack)
