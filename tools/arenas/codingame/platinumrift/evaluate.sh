#!/bin/bash

# Define some source locations
srcs="usersrc"
bins="userbin"

# Clean up previous leftovers
rm -r -f $srcs $bins
mkdir -p $srcs $bins

# Parametrize the input submission
./parametrize.sh "$@" build

# Iterate over all compiled binaries and generate the output
echo > res.log
for ai in `find $bins -type f`; do
	echo "./platinumrift -user $ai -players $2" >> res.log
	./platinumrift -user $ai -players $2 >> res.log
	echo >> res.log
done

for ai in `find $bins -type f`; do
  cat res.log | grep wins | grep $ai:
done
