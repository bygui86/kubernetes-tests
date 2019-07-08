#!/bin/sh
# HTTPIE infinite loop

HOST="$1"
PORT=$2
ENDPOINT="$3"

while [ true ]
do
	http --pretty none --print b $HOST:$PORT/$ENDPOINT
	sleep 1
done
