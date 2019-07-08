#!/bin/sh
# TCP listener


HOST="$1"
PORT=$2
MSG="$3"

echo $MSG | nc -v $HOST $PORT
