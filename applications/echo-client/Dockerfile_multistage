
FROM golang:1.12-stretch AS gobuilder

WORKDIR /go/src/github.com/bygui86/kubernetes-tests/applications/echo-client
COPY . .

RUN ["/bin/bash", "-c", "go get -v -d ."]
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/app .

# ---

FROM alpine

RUN apk update --no-cache
RUN apk add --no-cache bash
RUN apk add --no-cache curl

WORKDIR /usr/bin/
COPY --from=gobuilder /bin/app .
COPY ./health-check.sh .

EXPOSE 7001
EXPOSE 7090

USER 1001

ENTRYPOINT "/usr/bin/app"
