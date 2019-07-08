#!/bin/sh
# CURL infinite loop

HOST="$1"
PORT=$2
ENDPOINT="$3"

while [ true ]
do
	curl --connect-timeout 3 --max-time 3 $HOST:$PORT/$ENDPOINT
	sleep 1
done
