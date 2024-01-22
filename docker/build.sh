#!/bin/bash

# Build the docker image， from the dockerfile-env
docker build -t moredoc:latest .

# Run the docker image
docker run -d --name moredoc-mysql --net docker-network -p 127.0.0.1:33060:3306 -e MYSQL_ROOT_PASSWORD=root mysql:8.0.36
docker run -it --name moredoc-server --net docker-network -p 127.0.0.1:8880:8880 -v $(pwd):/home/moredoc moredoc:latest