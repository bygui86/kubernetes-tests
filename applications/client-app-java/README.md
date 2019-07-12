
# Client application - Java

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
	mvnw clean spring-boot:run
	```

* On docker
	```
	docker build . -t client-app-java:latest
	docker run -d --name client-app \
		-e SPRING_PROFILES_ACTIVE=docker \
		-e HEAP_SIZE=256M \
		-e META_SIZE=300M \
		-p 8080:8080 -p 8090:8090 \
		client-app
	docker logs client-app-java -f
	```

---

## API calls
* Insert
	```
	http POST :8080/apis/users email="matteo.baiguini@rabbit.com" name="Matteo Baiguini" age=33
	http POST :8080/apis/users email="john.doe@rabbit.com" name="John Doe" age=42
	http POST :8080/apis/users email="jane.doe@rabbit.com" name="Jane Doe" age=24
	http POST :8080/apis/users email="clint.eastwood@rabbit.com" name="Clint Eastwood" age=75
	```
* Get
	```
	http :8080/apis/users
	http :8080/apis/users/matteo.baiguini@rabbit.com
	```
* Update
	```
	USER_ID=$(http :8080/apis/users/clint.eastwood@rabbit.com | jq ".id")
	http PUT :8080/apis/users id=$USER_ID email="clint.eastwood@rabbit.com" name="Clint Eastwood" age=89
	```
* Delete
	```
	http DELETE :8080/apis/users/jane.doe@rabbit.com
	http DELETE :8080/apis/users
	```

---

## REST endpoints

* `GET /users` get all
* `GET /users/{email}` get by email
* `POST /users` insert new
* `PUT /users` update
* `DELETE /users` delete all
* `DELETE /users/{email}` delete by email
