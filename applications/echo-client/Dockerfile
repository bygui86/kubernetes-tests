
FROM alpine

RUN apk update --no-cache
RUN apk add --no-cache bash
RUN apk add --no-cache curl

WORKDIR /usr/bin/
COPY ./echo-client ./app
COPY ./health-check.sh .

EXPOSE 7001
EXPOSE 7090

USER 1001

ENTRYPOINT "/usr/bin/app"
