#!/bin/bash


CREATE_OPTS="--driver digitalocean --digitalocean-access-token $DIGITALOCEAN_ACCESS_TOKEN --digitalocean-image ubuntu-17-04-x64 --digitalocean-region fra1"


#CREATE_OPTS="--driver virtualbox --virtualbox-memory=1024 --engine-insecure-registry=192.168.99.104"

#
# create manager and init swarm
#
docker-machine create $CREATE_OPTS manager
eval $(docker-machine env manager)
docker swarm init --advertise-addr $(docker-machine ip manager) --listen-addr $(docker-machine ip manager):2377


JOIN_TOKEN=$(docker swarm join-token -q worker)

#
# create nodes
#
docker-machine create $CREATE_OPTS node-1 &
PID1=$!
docker-machine create $CREATE_OPTS node-2 &
PID2=$!
docker-machine create $CREATE_OPTS node-3 &
PID3=$!

wait $PID1
wait $PID2
wait $PID3

#
# join the swarm cluster
#
docker-machine ssh node-1 docker swarm join --token $JOIN_TOKEN $(docker-machine ip manager):2377
docker-machine ssh node-2 docker swarm join --token $JOIN_TOKEN $(docker-machine ip manager):2377
docker-machine ssh node-3 docker swarm join --token $JOIN_TOKEN $(docker-machine ip manager):2377

docker info | grep Nodes
