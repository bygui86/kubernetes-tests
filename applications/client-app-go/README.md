
# Client App Go

## Instructions

1. Prepare
	```shell
	mkdir -p $GOPATH/src/github.com
	cd $GOPATH/src/github.com
	git clone git@github.com:bygui86/kubernetes-tests.git
	cd bygui86/kubernetes-tests/applications/client-app-go
	go get ./...
	```

2. Build
	```shell
	go build .
	```

3. Build Docker image
	* from build
	```shell
	CGO_ENABLED=0 && GOOS=linux && go build -a -installsuffix cgo .
	docker build . -t client-app-go:latest
	```
	* multistage
	```shell
	docker build . -f Dockerfile_multistage -t client-app-go:latest
	```

4. Run
	* from code
	```shell
	go run main.go
	```
	* from compiled
	```shell
	./client-app-go
	```
	* as container
	```shell
	docker run -d --name client-app-go -p 8080:8080 -p 8090:8090 client-app-go:latest && docker logs client-app-go -f
	```

5. Test
	* POST
		```shell
		http POST :8080/apis/users email="matteo.baiguini@rabbit.com" name="Matteo Baiguini" age=33
		http POST :8080/apis/users email="john.doe@rabbit.com" name="John Doe" age=42
		http POST :8080/apis/users email="jane.doe@rabbit.com" name="Jane Doe" age=24
		http POST :8080/apis/users email="clint.eastwood@rabbit.com" name="Clint Eastwood" age=75
		```
	* GET
		```shell
		http :8080/apis/users
		http :8080/apis/users/matteo.baiguini@rabbit.com
		```
	* PUT
		```shell
		http PUT :8080/apis/users id=$(http :8080/apis/users/clint.eastwood@rabbit.com | jq ".id") email="clint.eastwood@rabbit.com" name="Clint Eastwood" age=89
		```
	* DELETE
		```shell
		http DELETE :8080/apis/users/clint.eastwood@rabbit.com
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

---

## Versions

### 0.0.1
- [x] rest apis
- [x] kubernetes probes
- [x] go structure
- [x] expose logs

### 0.0.2 - `TODO`
- [ ] expose metrics
- [ ] expose tracing

### 0.0.3 - `TODO`
- [ ] use kustomize for kube manifests

---

## Links

* [Tutorial](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql)
