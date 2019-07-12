# Kubernetes tests
Pool of simple tests to enhance understanding of Kubernetes, Istio and cloud world

## Sections

* [kubernetes](kubernetes)
* [applications](applications)
* [docker](docker)
* [istio](istio)

---

## TODOs

### version 1
- [x] server-app
- [x] client-app-java
- [x] client-app-go
- [x] postgres operator on kube
- [ ] logging on kube - `IN PROGRESS`
- [ ] monitoring on kube - `IN PROGRESS`
- [ ] tracing on kube
- [ ] konstallate on kube

### version 2
- [ ] traefik on kube
- [ ] rbac
- [ ] network limitations - `IN PROGRESS`
  - [ ] introduce calico
  - [ ] network policies

### version 3
- [ ] introduce spring-cloud-kubernetes in server-app
- [ ] introduce skaffold/jib
- [ ] introduce kustomize

### version 4
- [ ] istio
- [ ] flagger

### best effort
- [ ] install ci/cd on kube
- [ ] ambassador on kube
- [ ] server-app with java11 modules

---

## Kubernetes aspects

- [x] logging > EFK (elasticsearch, fluentd, kibana)
	- [x] manifests / operator
	- [ ] auto-config
- [ ] monitoring > Prometheus, Grafana
	- [x] manifests / operator
	- [x] auto-config
	- [x] prometheus-node-exporter
	- [ ] prometheus-postgres-exporter
	- [ ] influxdb as prometheus db
- [ ] tracing > Jaeger
	- [ ] manifests / operator
	- [ ] auto-config
- [ ] istio
	- [ ] manifests / operator
	- [ ] auto-config
- [ ] flagger
	- [ ] manifests / operator
	- [ ] auto-config
- [ ] ci/cd > JenkinsX, Prow, Tekton
	- [ ] manifests
	- [ ] auto-config

---

## Applications

- [x] server application
	- [x] logs
	- [x] metrics
	- [x] tracing
	- [x] docker multistage
	- [x] kubernetes manifests
	- [x] h2 version
	- [x] mongo version
	- [x] mysql version
	- [x] postgres version
- [x] client application java
	- [x] logs
	- [x] metrics
	- [x] tracing
	- [x] docker multistage
	- [x] kubernetes manifests
- [ ] client application golang
	- [x] logs
	- [ ] metrics
	- [ ] tracing
	- [x] docker multistage
	- [x] kubernetes manifests
- [x] log server
	- [x] logs
	- [x] docker multistage
	- [x] kubernetes manifests
	- [x] istio mirroring
- [x] echo server
	- [x] logs
	- [x] docker multistage
	- [x] kubernetes manifests
	- [x] kube shell probes
	- [x] istio testing
- [x] echo client
	- [x] logs
	- [x] docker multistage
	- [x] kubernetes manifests
	- [x] kube shell probes
	- [x] istio testing
- [x] sleepybox
	- [x] install networking tools
	- [x] docker multistage
	- [x] kubernetes manifests
	- [x] istio testing

---

## Links

### Kubernetes
* https://medium.com/deepaksood619/ultimate-kubernetes-infrastructure-monitoring-metrics-logs-c7b871d797bd

### CI/CD
* https://blog.csanchez.org/2019/03/05/progressive-delivery-with-jenkins-x-automatic-canary-deployments/

### Applications
* https://www.baeldung.com/spring-cloud-sleuth-single-application
* https://www.baeldung.com/spring-cloud-kubernetes
