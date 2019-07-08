#!/bin/sh
# TCP listener


PORT=8600

while [ true ]
do
	nc -l -p $PORT -v
	sleep 1
done
