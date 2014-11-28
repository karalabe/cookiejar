#!/bin/bash

# Parses a source file, retrieves all parametrized values and replaces them one
# by one with some predefined ranges and steps, compiling each version.

# Make sure output folders exist
srcs="usersrc"
bins="userbin"
mkdir -p $srcs $bins

# Parametrizes the decay argument
function phase_change {
	inject $1 $2 PHASE_CHANGE "${@:3}"
}

function phase_slope {
	inject $1 $2 PHASE_SLOPE "${@:3}"
}

function access_decay {
	inject $1 $2 ACCESS_DECAY "${@:3}"
}

function access_weight {
	inject $1 $2 ACCESS_WEIGHT "${@:3}"
}

function platinum_decay {
	inject $1 $2 PLATINUM_DECAY "${@:3}"
}

function territory_decay {
	inject $1 $2 TERRITORY_DECAY "${@:3}"
}

function presence_decay {
	inject $1 $2 PRESENCE_DECAY "${@:3}"
}

function threat_weight {
	inject $1 $2 THREAT_WEIGHT "${@:3}"
}

# Executes a parameter injection, cascading the function chain
function inject {
	# file=$1
	# name=$2
	# var=$3
	# vals=$4
	# next=$5

	IFS=' ' read -ra ARRAY <<< "$4"
	for PARAM in "${ARRAY[@]}"; do
		out=${2}_$PARAM
		sed "s|\(${3}_${players}P[ ]*=\)[ ]*[0-9\\.]*[ ]*//[ ]*@PARAM|\1 $PARAM|" < $1 >$srcs/$out.go
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
players=$2
func=$3

$func $file $name "${@:4}"
