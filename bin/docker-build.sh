#!/usr/bin/env bash

mm_wiki_dir=/home/ice/Storage/mm-wiki

dir_path=$(cd `dirname $0`; pwd)

cd ${dir_path}/../

docker rm mm-wiki
docker rmi mm-wiki

docker build -t mm-wiki .
docker run -d -p 8091:8091 --net mynetwork --ip 172.18.0.12 -v ${mm_wiki_dir}/conf/:/opt/mm-wiki/conf/ -v ${mm_wiki_dir}/data/:/data/mm-wiki/data/ --name mm-wiki mm-wiki

docker images | grep '<none>' | awk '{print $3}' | xargs docker rmi