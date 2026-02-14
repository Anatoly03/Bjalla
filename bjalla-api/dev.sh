#!/usr/bin/env bash

# Locate and navigate to  the directory of this start.sh file.
BASEDIR=$(dirname "$0")
cd $BASEDIR

# Serve the API module
go run src/main.go serve
