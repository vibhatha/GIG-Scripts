#!/bin/bash

# Check if the root directory is provided as an argument
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <rootDir>"
    exit 1
fi

rootDir=$1

# Pass the rootDir as an argument to the Go program
go run import_orgchart_data.go import_csv.go "$rootDir"