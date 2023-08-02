#!/bin/bash

# make sure the sqlite data file is there
if ![ -e "/data/agenda.sql" ]; then mkdir -p "/data"; touch "/data/agenda.sql"; fi

# build and run docker
docker build -t agenda.latest -f ./build/Dockerfile

is_running=$(docker ps | grep agenda | wc -l)

if [ "$is_running" != "1" ]; then docker stop agenda && docker rm agenda; fi

docker run -d --name=agenda -v /data:/app/data  --env_file ./env_vars agenda.latest