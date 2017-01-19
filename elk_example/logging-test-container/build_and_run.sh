#!/bin/bash

docker build -t logging-test-container .
docker run --rm --name logging-test-container --log-driver=gelf --log-opt gelf-address=udp://127.0.0.1:5000 logging-test-container

exit 0
