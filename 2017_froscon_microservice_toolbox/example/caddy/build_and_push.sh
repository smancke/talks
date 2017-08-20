#!/bin/bash

docker build -t demo-registry.mancke.io/caddy .
docker push demo-registry.mancke.io/caddy
