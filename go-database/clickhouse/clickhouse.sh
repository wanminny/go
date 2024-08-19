#!/bin/zsh

set -e -x

mkdir -p ~/clickhouse-data/{data,conf,log}
docker pull clickhouse/clickhouse-server:latest

docker run -d --rm --name clickhouse-server --ulimit nofile=262144:262144 clickhouse/clickhouse-server

docker cp clickhouse-server:/etc/clickhouse-server/config.xml ~/clickhouse-data/conf/config.xml
docker cp clickhouse-server:/etc/clickhouse-server/users.xml ~/clickhouse-data/conf/users.xml

docker stop clickhouse-server

docker run -d --name=clickhouse-server \
-p 8123:8123 -p 9090:9000 \
--ulimit nofile=262144:262144 \
-v ~/clickhouse-data/data:/var/lib/clickhouse:rw \
-v ~/clickhouse-data/conf/config.xml:/etc/clickhouse-server/config.xml \
-v ~/clickhouse-data/conf/users.xml:/etc/clickhouse-server/users.xml \
-v ~/clickhouse-data/log:/var/log/clickhouse-server:rw \
clickhouse/clickhouse-server