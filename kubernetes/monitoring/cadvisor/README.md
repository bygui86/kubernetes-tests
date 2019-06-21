# Kubernetes monitoring - cAdvisor

## Instructions

1. Clone repo
	```shell
	git clone git@github.com:google/cadvisor.git
	```

2. Generate base daemonset and deploy
   ```shell
   kustomize build deploy/kubernetes/base | kubectl apply -f -
   ```

---

## Links
* [repo](https://github.com/google/cadvisor)
* [deploy](https://github.com/google/cadvisor/tree/master/deploy/kubernetes)
