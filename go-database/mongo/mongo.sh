#!/bin/bash

docker pull mongo:4.0.4
mkdir -p ~/mongo-data
docker run -p 27017:27017 -v ~/mongo-data:/data/db --name mongodb -d mongo

