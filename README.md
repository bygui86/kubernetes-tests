# Kubernetes test applications
Pool of simple applications to enhance testing of Kubernetes various aspects 

---

## TODOs

### version 1
- [x] server-app
- [x] client-app-java
- [ ] postgres operator on kube
- [ ] logging on kube
- [ ] monitoring on kube
- [ ] tracing on kube
- [ ] mapping on kube
- [ ] traefik on kube
- [ ] konstallate on kube

### version 2
- [ ] introduce spring-cloud-kubernetes in server-app
- [ ] introduce skaffold/jib
- [ ] introduce kustomize
- [ ] install ci/cd on kube

### version 3
- [ ] server-app with java11 modules
- [ ] client-app-go
- [ ] rbac
- [ ] network limitations
  - [ ] introduce calico
  - [ ] network policies

### version 4
- [ ] istio

---

## Kubernetes aspects

- [x] logging > EFK (elasticsearch, fluentd, kibana)
	- [x] manifests / operator
	- [ ] auto-config
- [ ] monitoring > Prometheus, Grafana
	- [x] manifests / operator
	- [ ] auto-config
	- [ ] influxdb as prometheus db
	- [x] prometheus-node-exporter
	- [ ] prometheus-postgres-exporter
- [ ] tracing > Jaeger
	- [ ] manifests / operator
	- [ ] auto-config
- [ ] mapping > Kiali
	- [ ] manifests / operator
	- [ ] auto-config
- [ ] ci/cd > JenkinsX, Prow, Tekton
	- [ ] manifests
	- [ ] auto-config

---

## Applications

- [ ] server application
	- [x] logs
	- [x] metrics
	- [x] tracing
	- [x] docker multistage
	- [ ] kubernetes manifests
- [ ] client application java
	- [ ] logs
	- [ ] metrics
	- [ ] tracing
	- [ ] docker multistage
	- [ ] kubernetes manifests
- [ ] client application golang
	- [ ] logs
	- [ ] metrics
	- [ ] tracing
	- [ ] docker multistage
	- [ ] kubernetes manifests

---

## Links

### Kubernetes
* https://medium.com/deepaksood619/ultimate-kubernetes-infrastructure-monitoring-metrics-logs-c7b871d797bd

### CI/CD
* https://blog.csanchez.org/2019/03/05/progressive-delivery-with-jenkins-x-automatic-canary-deployments/

### Applications
* https://www.baeldung.com/spring-cloud-sleuth-single-application
* https://www.baeldung.com/spring-cloud-kubernetes
