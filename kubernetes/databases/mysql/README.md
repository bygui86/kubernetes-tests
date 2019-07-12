
# Kubernetes databases - MySQL

## Instructions

### Stand-alone

1. Start minikube
	```shell
	minikube start \
		--profile mysql \
		--vm-driver hyperkit --cpus 4 --memory 12288 \
		--enable-default-cni --network-plugin=cni \
		--extra-config=apiserver.authorization-mode=RBAC
	```

2. Deploy
	```shell
	helm install --name server -f values.yaml stable/mysql
	```

### Operator

`TODO`

---

## Links
