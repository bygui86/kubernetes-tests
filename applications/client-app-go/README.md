
# Client App Go

## Build & run

* Locally
	```
	go build .
	go run main.go
		or
	./client-app-go
	```

* On docker
	```
	docker build . -t client-app-go:latest
	docker run -d --name client-app-go -p 8080:8080 -p 8090:8090 client-app-go:latest
	docker logs client-app-go -f
	```

---

## API calls
* POST
	```
	http POST :8080/apis/users email="matteo.baiguini@rabbit.com" name="Matteo Baiguini" age=33
	http POST :8080/apis/users email="john.doe@rabbit.com" name="John Doe" age=42
	http POST :8080/apis/users email="jane.doe@rabbit.com" name="Jane Doe" age=24
	http POST :8080/apis/users email="clint.eastwood@rabbit.com" name="Clint Eastwood" age=75
	```
* GET
	```
	http :8080/apis/users
	http :8080/apis/users/matteo.baiguini@rabbit.com
	```
* PUT
	```
	http PUT :8080/apis/users id=$(http :8080/apis/users/clint.eastwood@rabbit.com | jq ".id") email="clint.eastwood@rabbit.com" name="Clint Eastwood" age=89
	```
* DELETE
	```
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
- [ ] expose traces

### 0.0.3 - `TODO`
- [ ] use kustomize for kube manifests

---

## Links

* [Tutorial](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql)
