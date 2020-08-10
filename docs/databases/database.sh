#!/usr/bin/env bash

dir=$(cd `dirname $0`; pwd)
cd $dir

mysql -h127.0.0.1 -uroot -p123456 -e "create database mm_wiki"
mysql -h127.0.0.1 -uroot -p123456 mm_wiki < ./table.sql
mysql -h127.0.0.1 -uroot -p123456 mm_wiki < ./data.sql
