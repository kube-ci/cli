#!/usr/bin/env bash

pushd $GOPATH/src/github.com/kube-ci/cli/hack/gendocs
go run main.go
popd
