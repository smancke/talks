#!/bin/bash

docker build -t demo-registry.mancke.io/content .
docker push demo-registry.mancke.io/content
