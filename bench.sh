#!/bin/bash

# Change to the project root
cd `dirname $0`

# Run all the tests
echo "Running tests..."
go test ./...

# Run all the benchmarks, individually (race condition bug in go test)
cpu=`cat /proc/cpuinfo | grep "model name" | head -n 1 | cut -d ':' -f 2 | sed -e 's/^ *//g' -e 's/  */ /g' -e 's/ *$//g'`
echo
echo "Bencmark results on $cpu:"

packs=`find . -mindepth 1 -type d | grep -v "\./\." | sort`
for pkg in $packs; do
  echo "- `basename $pkg`"
  go test -run=NONE -bench=. -benchtime=100ms $pkg | grep Benchmark | awk '{print "    -", $1, "\t", $3, "\t", $4}'
done
