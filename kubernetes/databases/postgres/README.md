
# Kubernetes databases - PostgreSQL

## Instructions

### Stand-alone

1. Start minikube
	```
	minikube start \
		--profile postgres \
		--vm-driver hyperkit --cpus 4 --memory 12288 \
		--enable-default-cni --network-plugin=cni \
		--extra-config=apiserver.authorization-mode=RBAC
	```

2. Deploy
	* automatic using Helm
		```
		helm install --name postgres -f values.yaml stable/postgresql
		```
	* manual
		```
		helm template --name postgresql -f values.yaml . > postgresql.yaml
		kubectl apply -f postgresql.yaml
		```

### Cluster

`TODO`

### Operator

#### Zalando

1. Start minikube
	```
	minikube start \
		--profile postgres \
		--vm-driver hyperkit --cpus 4 --memory 12288 \
		--enable-default-cni --network-plugin=cni \
		--extra-config=apiserver.authorization-mode=RBAC
	```

2. Clone repo
	```
	git clone https://github.com/zalando/postgres-operator.git
	cd postgres-operator
	```

3. Deploy operator
	* Manually
		```
		kubectl apply -f manifests/configmap.yaml
		kubectl apply -f manifests/operator-service-account-rbac.yaml
		kubectl apply -f manifests/postgres-operator.yaml
		kubectl get pod -l name=postgres-operator -w
		```
	* Helm
		```
		helm init
		helm install --name zalando ./charts/postgres-operator
		kubectl get pod -l app.kubernetes.io/name=postgres-operator
		```

4. Spin up Postgres
	```
	kubectl apply -f manifests/minimal-postgres-manifest.yaml
	```

5. Check
	* Postgres cluster
		```
		kubectl get postgresql
		```
	* Pods
		```
		kubectl get pods -l application=spilo -L spilo-role
		```
	* Services
		```
		kubectl get svc -l application=spilo -L spilo-role
		```

6. Get credentials
	```
	echo "Postgres HOST: " $(minikube service acid-minimal-cluster --url | sed 's,.*/,,' | cut -d: -f 1)
	echo "Postgres PORT: " $(minikube service acid-minimal-cluster --url | sed 's,.*/,,' | cut -d: -f 2)
	echo "Postgres USERNAME: " $(kubectl get secret postgres.acid-minimal-cluster.credentials -o 'jsonpath={.data.username}' | base64 -D)
	echo "Postgres PASSWORD: " $(kubectl get secret postgres.acid-minimal-cluster.credentials -o 'jsonpath={.data.password}' | base64 -D)
	```

`WARN: Kubernetes service manifests are not 100% correct. Moreover they are compliant with Istio (in some service definitions miss the 'selector' field)`
7. Fix service definition
	```
	kubectl apply -f postgres-operator-svc.yaml
	```


#### CrunchyData

`TODO`
`INFO: Seems to be more complicated than Zalando one`

---

## Links
* Zalando
  * https://postgres-operator.readthedocs.io/en/latest/quickstart/
  * https://github.com/zalando/postgres-operator
* CrunchyData
  * https://github.com/CrunchyData/postgres-operator
  * https://access.crunchydata.com/documentation/postgres-operator/4.0.0/installation/operator-install/
