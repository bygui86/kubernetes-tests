
FROM alpine

RUN apk update --no-cache
RUN apk add --no-cache bash
RUN apk add --no-cache curl

WORKDIR /usr/bin/
COPY ./client-app-go ./app
COPY ./health-check.sh .

EXPOSE 8080
EXPOSE 8090

USER 1001

ENTRYPOINT "/usr/bin/app"
