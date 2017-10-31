#!/bin/bash



go build .
docker build -t edge_service .

