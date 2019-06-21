
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

## Build

### Classic
```bash
mvnw clean package
docker build . -t server-app:latest
```

### Multistage
```bash
docker build . -f Dockerfile_multistage -t server-app:latest
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
kubectl apply -f kube/configmap.yaml
kubectl apply -f kube/secret.yaml
kubectl apply -f kube/service.yaml
kubectl apply -f kube/deployment.yaml
kubectl apply -f kube/hor-pod-scale.yaml
```

## API calls

### Local / Docker
### Insert
```bash
http POST :8080/users email="matteo.baiguini@rabbit.com" name="Matteo Baiguini" age=33
http POST :8080/users email="john.doe@rabbit.com" name="John Doe" age=42
http POST :8080/users email="jane.doe@rabbit.com" name="Jane Doe" age=24
http POST :8080/users email="clint.eastwood@rabbit.com" name="Clint Eastwood" age=75
```
### Get
```bash
http :8080/users
http :8080/users/matteo.baiguini@rabbit.com
```
### Update
```bash
USER_ID=$(http :8080/users/clint.eastwood@rabbit.com | jq ".id")
http PUT :8080/users id=$USER_ID email="clint.eastwood@rabbit.com" name="Clint Eastwood" age=89
```
### Delete
```bash
http DELETE :8080/users/jane.doe@rabbit.com
http DELETE :8080/users
```

### Minikube
### 
```bash
SERVER_APP_HOST=$(minikube ip)
SERVER_APP_PORT=$(k get svc | grep -i server-app | awk '{print $5}' | cut -d ',' -f 1 | sed 's,8080:,,' | sed 's,/TCP,,')
```
### Insert
```bash
http POST $SERVER_APP_HOST:$SERVER_APP_PORT/users email="matteo.baiguini@rabbit.com" name="Matteo Baiguini" age=33
http POST $SERVER_APP_HOST:$SERVER_APP_PORT/users email="john.doe@rabbit.com" name="John Doe" age=42
http POST $SERVER_APP_HOST:$SERVER_APP_PORT/users email="jane.doe@rabbit.com" name="Jane Doe" age=24
http POST $SERVER_APP_HOST:$SERVER_APP_PORT/users email="clint.eastwood@rabbit.com" name="Clint Eastwood" age=75
```
### Get
```bash
http $SERVER_APP_HOST:$SERVER_APP_PORT/users
http $SERVER_APP_HOST:$SERVER_APP_PORT/users/matteo.baiguini@rabbit.com
```
### Update
```bash
USER_ID=$(http $SERVER_APP_HOST:$SERVER_APP_PORT/users/clint.eastwood@rabbit.com | jq ".id")
http PUT $SERVER_APP_HOST:$SERVER_APP_PORT/users id=$USER_ID email="clint.eastwood@rabbit.com" name="Clint Eastwood" age=89
```
### Delete
```bash
http DELETE $SERVER_APP_HOST:$SERVER_APP_PORT/users/jane.doe@rabbit.com
http DELETE $SERVER_APP_HOST:$SERVER_APP_PORT/users
```
