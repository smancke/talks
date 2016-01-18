#!/bin/bash -e

export GOPATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

go get github.com/golang/mock/gomock &&\
go get github.com/golang/mock/mockgen &&\

set -x

$GOPATH/bin/mockgen -self_package objects -package objects \
            -destination $GOPATH/src/objects/mocks_test.go \
            objects \
            Flyable

