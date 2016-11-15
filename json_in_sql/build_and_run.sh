#!/bin/bash

cd `dirname $BASH_SOURCE`
export GOPATH=`pwd`
go get --tags json1 example
bin/example
