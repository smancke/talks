#!/bin/bash


CREATE_OPTS="--driver digitalocean --digitalocean-access-token $DIGITALOCEAN_ACCESS_TOKEN --digitalocean-image ubuntu-17-04-x64 --digitalocean-region fra1"
WORKER_COUNT=3


docker-machine create $CREATE_OPTS manager
eval $(docker-machine env manager)
docker swarm init --advertise-addr $(docker-machine ip manager) --listen-addr $(docker-machine ip manager):2377

JOIN_TOKEN=$(docker swarm join-token -q worker)

docker-machine create $CREATE_OPTS node-1 &
PID1=$!
docker-machine create $CREATE_OPTS node-2 &
PID2=$!
docker-machine create $CREATE_OPTS node-3 &
PID3=$!

wait $PID1
wait $PID2
wait $PID3

docker-machine ssh node-1 docker swarm join --token $JOIN_TOKEN $(docker-machine ip manager):2377
docker-machine ssh node-2 docker swarm join --token $JOIN_TOKEN $(docker-machine ip manager):2377
docker-machine ssh node-3 docker swarm join --token $JOIN_TOKEN $(docker-machine ip manager):2377

docker info | grep Nodes
