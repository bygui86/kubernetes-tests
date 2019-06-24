#!/bin/sh
# Kubernetes health check

set -e

HOST="$1"
PORT=$2
PATH="$3"

if [[ $(/usr/bin/curl --connect-timeout 2 --max-time 2 --write-out %{http_code} --silent --output /dev/null $HOST:$PORT/$PATH) == "200" ]]; then
	echo "KO :("
	exit -1
else
	echo "OK :)"
fi
