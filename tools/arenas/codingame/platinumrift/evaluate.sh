#!/bin/bash

echo >> res.log
echo "Starting evaluation" >> res.log

for ai in `find . | grep user_`; do
	echo "./platinumrift -user $ai" >> res.log
	./platinumrift -user $ai >> res.log
	echo >> res.log
done
