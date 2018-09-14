#!/bin/bash

export PATH="/usr/lib/go-1.10/bin:${PATH}"
export CGO_ENABLED=0
export GOOS="linux"
export GOPATH=`pwd`
export GOBIN="${GOPATH}/bin"
export GOEXE="autoservice"

go build -o bin/wxtaoke/wxtaoke game/main/wxtaoke

