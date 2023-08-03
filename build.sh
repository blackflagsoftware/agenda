#!/bin/bash

# make sure the sqlite data file is there
if [ ! -e "/Users/ericsmith/data/agenda.sql" ]; then mkdir -p "/Users/ericsmith/data"; touch "/Users/ericsmith/data/agenda.sql"; fi

# build and run docker
docker build -t agenda:latest -f ./build/Dockerfile .
is_running=$(docker ps | grep agenda | wc -l | xargs)
is_container=$(docker ps -a | grep agenda | wc -l | xargs)

if [ "$is_running" == "1" ]; then docker stop agenda; fi 
if [ "$is_container" == "1" ]; then docker rm agenda; fi

docker run -d --name=agenda -v /Users/ericsmith/data:/app/data -p 12580:12580 --env-file ./env_vars agenda:latest