#!/bin/bash

PORT=""

while getopts "p:h" opt; do
    case $opt in
        p)
            PORT=":$OPTARG"
            echo "Port specified: $OPTARG" >&2
            ;;
    esac
done


go run $PWD/cmd/api -port=$PORT
