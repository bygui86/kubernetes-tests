
FROM alpine:latest

RUN apk add --update curl
RUN apk add --no-cache bash
RUN apk add --no-cache util-linux
RUN rm -rf /var/cache/apk/*

WORKDIR /usr/bin/
COPY entrypoint.sh .
COPY ./storage-filler.sh .
RUN chmod +x storage-filler.sh

ENTRYPOINT ["/entrypoint.sh"]
CMD ["curl"]
