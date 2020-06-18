#!/usr/bin/env bash

# 挂载路径
mm_wiki_dir=/home/ice/Storage/mm-wiki
# 脚本路径
dir_path=$(cd `dirname $0`; pwd)
# 项目路径
project_dir=${dir_path}/../
# 打包路径
dist_path=${project_dir}/dist

\rm -rf ${dist_path}
mkdir -p ${dist_path}

#先创建目录
mkdir -p ${dist_path}/mm-wiki
#将项目里用到的文件目录copy过来
cd ${project_dir}
cp -r app conf docs global install scripts static views vendor LICENSE README.md build.sh go.mod go.sum main.go pack.sh router.go ${dist_path}/mm-wiki/
cp docker/Dockerfile-local ${dist_path}/
# 并打包
cd ${dist_path}
tar -zcvf mm-wiki.tar.gz mm-wiki

# 停止删除原来的容器
docker stop mm-wiki
docker rm mm-wiki
docker rmi mm-wiki

docker build -t mm-wiki . -f Dockerfile-local
docker run -d -p 8091:8091 --net mynetwork --ip 172.18.0.12 -v ${mm_wiki_dir}/conf/:/opt/mm-wiki/conf/ -v ${mm_wiki_dir}/data/:/data/mm-wiki/data/ --name mm-wiki --restart=always mm-wiki

#删除过程中的无用容器
docker images | grep '<none>' | awk '{print $3}' | xargs docker rmi
