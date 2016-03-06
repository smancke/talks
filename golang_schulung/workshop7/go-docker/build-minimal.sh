#!/bin/bash

CGO_ENABLED=0 go build google_query.go
cp /etc/ssl/certs/ca-certificates.crt .
docker build -f Dockerfile.minimal -t google-query-minimal .
