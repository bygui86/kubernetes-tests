
# docker build . -t sleepybox:latest
# docker run -d --rm --name sleepybox sleepybox:latest

FROM alpine:latest

RUN apk update --no-cache
RUN apk add --no-cache bash
RUN apk add --no-cache curl
RUN apk add --no-cache netcat-openbsd
RUN apk add --no-cache ospd-netstat
RUN apk add --no-cache iptables
RUN apk add --no-cache tcpdump
RUN apk add --no-cache screen
RUN apk add --no-cache pv
RUN rm -rf /var/cache/apk/*

ENTRYPOINT ["/bin/sh", "-c", "while :; do echo 'I will go to sleep for an hour, see ya later...'; sleep 3600; done"]
