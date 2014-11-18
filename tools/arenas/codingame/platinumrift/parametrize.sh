#!/bin/bash

# Parses a source file, retrieves all parametrized values and replaces them one
# by one with some predefined ranges and steps, compiling each version.

# Make sure output folders exist
srcs="usersrc"
bins="userbin"
mkdir -p $srcs $bins

# Parametrizes the decay argument
function decay {
	inject DECAY_RATE "1 2 3 4 5 6 7 8 9 10" "$@"
}

function lmfao {
	inject LMFAO "1.8 2.8 3.8 4.8 5.8 6.8 7.8 8.8 9.8 10.8" "$@"
}

# Executes a parameter injection, cascading the function chain
function inject {
	# var=$1
	# vals=$2
	# file=$3
	# name=$4
	# next=$5

	IFS=' ' read -ra ARRAY <<< "$2"
	for PARAM in "${ARRAY[@]}"; do
		out=${4}_$PARAM
		sed "s|\($1[ ]*=\)[ ]*[0-9]*[ ]*//[ ]*@PARAM|\1 $PARAM|" < $3 >$srcs/$out.go
		$5 $srcs/$out.go $out "${@:6}"
	done
}

# Builds the final binary for evaluation
function build {
	echo "Building $2..."
	go build -o $bins/$2 $1
}

# Fetch the base name, source file and chain through the parametrizers
file=$1
name="${file%.*}"
func=$2

$func $file $name "${@:3}"
