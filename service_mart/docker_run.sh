#!/usr/bin/env bash
docker run --name redis -p 127.0.0.1:6379:6379 -d daocloud.io/library/redis:3.2.9
docker run --name mysql -p 127.0.0.1:3306:3306 -v `替换为本地数据路径`:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root  -d daocloud.io/library/mysql:5.7
docker build -t mart-service .
docker run --name mart-service -it -v $(pwd):/app --link redis:redis --link mysql:mart-host -p 127.0.0.1:8000:8000 -d mart-service

# 数据库创建用户新用户，并赋予权限
#grant all privileges on *.* to sg@'%' identified by 'sg123456';
#flush privileges;

