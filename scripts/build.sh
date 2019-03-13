#!/bin/bash

set -e

CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o next .