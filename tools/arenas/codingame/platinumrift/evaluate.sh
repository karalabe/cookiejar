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
for ai in `find $bins -type f`; do
	echo "./platinumrift -user $ai" >> res.log
	./platinumrift -user $ai >> res.log
	echo >> res.log
done
