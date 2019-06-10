
# Client application - Java

## TODOs

- [x] Implementation
- [x] Metrics
- [x] Tracing
- [x] Swagger
- [x] Docker multistage
- [ ] Kubernetes
	- [ ] implementation
		- [ ] service-discovery
		- [ ] configmap
		- [ ] secrets
	- [ ] manifests

---

### Prerequisites
* Docker
* Minikube / Cloud Kubernetes cluster
* httpie
* jq

---

## Build

### Multistage
```bash
docker build . -t client-app-java:latest
```

### Classic
```bash
mvnw clean package
docker build . -f Dockerfile_classic -t client-app-java:latest
```

## Run

### Locally
```bash
mvnw clean spring-boot:run
```

### Docker
```bash
docker run -d --name client-app \
	-e SPRING_PROFILES_ACTIVE=docker \
	-e HEAP_SIZE=256M \
	-e META_SIZE=300M \
	-p 8080:8080 -p 8090:8090 \
	client-app
```

### Kubernetes
```bash
TODO
```

## API calls 

### Insert
```bash
http POST :8080/apis/users email="matteo.baiguini@rabbit.com" name="Matteo Baiguini" age=33
http POST :8080/apis/users email="john.doe@rabbit.com" name="John Doe" age=42
http POST :8080/apis/users email="jane.doe@rabbit.com" name="Jane Doe" age=24
http POST :8080/apis/users email="clint.eastwood@rabbit.com" name="Clint Eastwood" age=75
```

## Get
```bash
http :8080/apis/users
http :8080/apis/users/matteo.baiguini@rabbit.com
```

## Update
```bash
USER_ID=$(http :8080/apis/users/clint.eastwood@rabbit.com | jq ".id")
http PUT :8080/apis/users id=$USER_ID email="clint.eastwood@rabbit.com" name="Clint Eastwood" age=89
```

## Delete
```bash
http DELETE :8080/apis/users/jane.doe@rabbit.com
http DELETE :8080/apis/users
```
