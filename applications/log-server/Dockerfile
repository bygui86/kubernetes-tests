
FROM alpine

RUN apk update --no-cache
RUN apk add --no-cache bash
RUN apk add --no-cache curl

WORKDIR /usr/bin/
COPY ./log-server ./app

EXPOSE 7001
EXPOSE 7090

USER 1001

ENTRYPOINT "/usr/bin/app"
