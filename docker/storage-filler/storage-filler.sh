#!/bin/bash
# STORAGE filler

set -o errexit -o nounset -o pipefail

COUNTER=1
while [ true ]
do
	FILE_NAME="data/filler-$COUNTER.tmp"
	if [[ ! -f $FILE_NAME ]]; then
		echo "Creating $FILE_NAME"
		fallocate -l 100M $FILE_NAME
	else
		echo "$FILE_NAME already exists"
		COUNTER=$((COUNTER+100))
	fi
	COUNTER=$((COUNTER+1))
	df -h | grep -i '/usr/bin/data'
	sleep 10
done
