#!/bin/bash -e

cd `dirname $BASH_SOURCE`

docker-compose up -d

export GOPATH=`pwd`
go get --tags json1 example
bin/example
