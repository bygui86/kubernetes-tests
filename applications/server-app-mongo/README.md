
# Server application

## TODOs

- [x] Implementation
- [x] Metrics
- [x] Tracing
- [x] Swagger
- [x] Docker multistage
- [ ] Kubernetes
	- [x] manifests
	- [ ] implementation
		- [ ] service-discovery
		- [ ] configmap
		- [ ] secrets

---

### Prerequisites
* Docker
* Minikube / Cloud Kubernetes cluster
* httpie
* jq

---

## Build & run

* Locally
	```
	docker run -d --name mongo \
		-p 27017:27017 \
		mongo
	mvnw clean spring-boot:run
	```

* On Docker
	```
	docker build . -t server-app-mongo:latest
	docker network create test
	docker run -d --name mongo \
		-p 27017:27017 \
		--network test \
		mongo
	docker run -d --name server-app-mongo \
		-e SPRING_PROFILES_ACTIVE=docker \
		-e HEAP_SIZE=256M \
		-e META_SIZE=300M \
		-p 8080:8080 -p 8090:8090 \
		--network test \
		server-app-mongo
	docker logs server-app-mongo -f
	```

* On Kubernetes
	```
	kubectl apply -f kube/configmap.yaml
	kubectl apply -f kube/secret.yaml
	kubectl apply -f kube/service.yaml
	kubectl apply -f kube/deployment.yaml
	kubectl apply -f kube/hor-pod-scale.yaml
	```

---

## API calls

### Local / Docker
* Insert
	```
	http POST :8080/users email="matteo.baiguini@rabbit.com" name="Matteo Baiguini" age=33
	http POST :8080/users email="john.doe@rabbit.com" name="John Doe" age=42
	http POST :8080/users email="jane.doe@rabbit.com" name="Jane Doe" age=24
	http POST :8080/users email="clint.eastwood@rabbit.com" name="Clint Eastwood" age=75
	```
* Get
	```
	http :8080/users
	http :8080/users/matteo.baiguini@rabbit.com
	```
* Update
	```
	USER_ID=$(http :8080/users/clint.eastwood@rabbit.com | jq ".id")
	http PUT :8080/users id=$USER_ID email="clint.eastwood@rabbit.com" name="Clint Eastwood" age=89
	```
* Delete
	```
	http DELETE :8080/users/jane.doe@rabbit.com
	http DELETE :8080/users
	```

### Minikube
* Prepare env-vars
	```
	SERVER_APP_HOST=$(minikube ip)
	SERVER_APP_PORT=$(k get svc | grep -i server-app | awk '{print $5}' | cut -d ',' -f 1 | sed 's,8080:,,' | sed 's,/TCP,,')
	```
* Insert
	```
	http POST $SERVER_APP_HOST:$SERVER_APP_PORT/users email="matteo.baiguini@rabbit.com" name="Matteo Baiguini" age=33
	http POST $SERVER_APP_HOST:$SERVER_APP_PORT/users email="john.doe@rabbit.com" name="John Doe" age=42
	http POST $SERVER_APP_HOST:$SERVER_APP_PORT/users email="jane.doe@rabbit.com" name="Jane Doe" age=24
	http POST $SERVER_APP_HOST:$SERVER_APP_PORT/users email="clint.eastwood@rabbit.com" name="Clint Eastwood" age=75
	```
* Get
	```
	http $SERVER_APP_HOST:$SERVER_APP_PORT/users
	http $SERVER_APP_HOST:$SERVER_APP_PORT/users/matteo.baiguini@rabbit.com
	```
* Update
	```
	USER_ID=$(http $SERVER_APP_HOST:$SERVER_APP_PORT/users/clint.eastwood@rabbit.com | jq ".id")
	http PUT $SERVER_APP_HOST:$SERVER_APP_PORT/users id=$USER_ID email="clint.eastwood@rabbit.com" name="Clint Eastwood" age=89
	```
* Delete
	```
	http DELETE $SERVER_APP_HOST:$SERVER_APP_PORT/users/jane.doe@rabbit.com
	http DELETE $SERVER_APP_HOST:$SERVER_APP_PORT/users
	```
