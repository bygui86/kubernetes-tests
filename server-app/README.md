
# Server application

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
docker build . -t server-app:latest
```

### Classic
```bash
mvnw clean package
docker build . -f Dockerfile_classic -t server-app:latest
```

## Run

### Locally
```bash
docker run -d --name postgres \
	-e POSTGRES_DB=server-app \
	-e POSTGRES_USER=user \
	-e POSTGRES_PASSWORD=secret \
	-p 5432:5432 \
	ostgres:alpine
mvnw clean spring-boot:run
```

### Docker
```bash
docker network create test
docker run -d --name postgres \
	-e POSTGRES_DB=server-app \
	-e POSTGRES_USER=user \
	-e POSTGRES_PASSWORD=secret \
	-p 5432:5432 \
	--network test \
	postgres:alpine
docker run -d --name server-app \
	-e SPRING_PROFILES_ACTIVE=docker \
	-e HEAP_SIZE=256M \
	-e META_SIZE=300M \
	-p 8080:8080 -p 8090:8090 \
	--network test \
	server-app
```

### Kubernetes
```bash
TODO
```

## API calls 

### Insert
```bash
http POST :8080/users email="matteo.baiguini@rabbit.com" name="Matteo Baiguini" age=33
http POST :8080/users email="john.doe@rabbit.com" name="John Doe" age=42
http POST :8080/users email="jane.doe@rabbit.com" name="Jane Doe" age=24
http POST :8080/users email="clint.eastwood@rabbit.com" name="Clint Eastwood" age=75
```

## Get
```bash
http :8080/users
http :8080/users/matteo.baiguini@rabbit.com
```

## Update
```bash
USER_ID=$(http :8080/users/clint.eastwood@rabbit.com | jq ".id")
http PUT :8080/users id=$USER_ID email="clint.eastwood@rabbit.com" name="Clint Eastwood" age=89
```

## Delete
```bash
http DELETE :8080/users/jane.doe@rabbit.com
http DELETE :8080/users
```
