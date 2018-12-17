#!/bin/bash
set -xeou pipefail

GOPATH=$(go env GOPATH)
REPO_ROOT="$GOPATH/src/github.com/kube-ci/cli"

export APPSCODE_ENV=prod

pushd $REPO_ROOT

rm -rf dist

./hack/make.py build
./hack/make.py push

rm -rf dist/.tag

popd
