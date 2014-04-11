#!/bin/sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
go run src/server.go src/view.go src/model.go