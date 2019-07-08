# Kubernetes tests - Applications - Sleepybox

## description
Simple Docker image containing networking tools

## ports
* 8600		TCP listener

## working directory
`/usr/bin`

## available scripts
* curl-inf-loop.sh $HOST $PORT $ENDPOINT

  to perform a rest calls to a HOST on a PORT at an ENDPOINT using curl

* http-inf-loop.sh $HOST $PORT $ENDPOINT

  to perform a rest calls to a HOST on a PORT at an ENDPOINT using httpie

* tcp-listener.sh

  to open a TCP listener on port 8600

* tcp-caller.sh $HOST $PORT $MSG

  to send a message (MSG) using TCP to a HOST on a PORT
