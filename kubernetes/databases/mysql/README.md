
# Kubernetes databases - MySQL

## Instructions

### Stand-alone

1. Start minikube
	```
	minikube start \
		--profile mysql \
		--vm-driver hyperkit --cpus 4 --memory 12288 \
		--enable-default-cni --network-plugin=cni \
		--extra-config=apiserver.authorization-mode=RBAC
	```

2. Deploy
	* automatic using Helm
		```
		helm install --name server -f values.yaml stable/mysql
		```
	* manual
		```
		helm template --name mysql -f values.yaml . > mysql.yaml
		kubectl apply -f mysql.yaml
		```

### Cluster

`TODO`

### Operator

`TODO`

---

## Links
