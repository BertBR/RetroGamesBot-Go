#! /bin/bash

echo "- Checking air..."
if ! command -v air &> /dev/null
then
    echo "MISSING DEPENDENCY: air is not installed"
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
    air -v
fi
