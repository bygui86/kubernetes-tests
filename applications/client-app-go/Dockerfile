
FROM alpine

RUN apk add --no-cache bash

WORKDIR /bin/
COPY ./client-app-go /bin/app

EXPOSE 8080
EXPOSE 8090

ENTRYPOINT "/bin/app"
