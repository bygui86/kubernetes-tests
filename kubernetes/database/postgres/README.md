
# Kubernetes databases - PostgreSQL

## Instructions

### Zalando operator

1. Start minikube
	```shell
	minikube start \
		--profile postgres \
		--vm-driver hyperkit --cpus 4 --memory 12288 \
		--enable-default-cni --network-plugin=cni \
		--extra-config=apiserver.authorization-mode=RBAC
	```

2. Clone repo
	```shell
	git clone https://github.com/zalando/postgres-operator.git
	cd postgres-operator
	```

3. Deploy operator
	* Manually
		```shell
		kubectl apply -f manifests/configmap.yaml
		kubectl apply -f manifests/operator-service-account-rbac.yaml
		kubectl apply -f manifests/postgres-operator.yaml
		kubectl get pod -l name=postgres-operator -w
		```
	* Helm
		```shell
		helm init
		helm install --name zalando ./charts/postgres-operator
		kubectl get pod -l app.kubernetes.io/name=postgres-operator
		```

4. Spin up Postgres
	```shell
	kubectl apply -f manifests/minimal-postgres-manifest.yaml
	```

5. Check
	* Postgres cluster
		```shell
		kubectl get postgresql
		```
	* Pods
		```shell
		kubectl get pods -l application=spilo -L spilo-role
		```
	* Services
		```shell
		kubectl get svc -l application=spilo -L spilo-role
		```

6. Get credentials
	```shell
	echo "Postgres HOST: " $(minikube service acid-minimal-cluster --url | sed 's,.*/,,' | cut -d: -f 1)
	echo "Postgres PORT: " $(minikube service acid-minimal-cluster --url | sed 's,.*/,,' | cut -d: -f 2)
	echo "Postgres USERNAME: " $(kubectl get secret postgres.acid-minimal-cluster.credentials -o 'jsonpath={.data.username}' | base64 -D)
	echo "Postgres PASSWORD: " $(kubectl get secret postgres.acid-minimal-cluster.credentials -o 'jsonpath={.data.password}' | base64 -D)
	```

### CrunchyData operator

`TODO`

---

## Links
* Zalando
  * https://postgres-operator.readthedocs.io/en/latest/quickstart/
  * https://github.com/zalando/postgres-operator
* CrunchyData
  * https://github.com/CrunchyData/postgres-operator
  * https://access.crunchydata.com/documentation/postgres-operator/4.0.0/installation/operator-install/
