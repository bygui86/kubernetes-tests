
# Kubernetes databases - MongoDB

## Instructions

### Stand-alone

1. Start minikube
	```
	minikube start \
		--profile mongo \
		--vm-driver hyperkit --cpus 4 --memory 12288 \
		--enable-default-cni --network-plugin=cni \
		--extra-config=apiserver.authorization-mode=RBAC
	```

2. Deploy
	* automatic using Helm
		```
		helm install --name mongodb -f values.yaml stable/mongodb
		```
	* manual
		```
		helm template --name mongodb -f values.yaml . > mongodb.yaml
		kubectl apply -f mongodb.yaml
		```

### Cluster (Replicaset in MongoDB naming convention)

1. Start minikube
	```
	minikube start \
		--profile mongo \
		--vm-driver hyperkit --cpus 4 --memory 12288 \
		--enable-default-cni --network-plugin=cni \
		--extra-config=apiserver.authorization-mode=RBAC
	```

2. Deploy
	* automatic using Helm
		```
		helm install --name server -f values.yaml stable/mongodb-replicaset
		```
	* manual
		`TODO`

### Operator

`TODO`

---

## Links
