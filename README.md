# Getting Started with Create Go App
 Have to set db first


# below record how to run this app at linux
sudo apt-get install docker.io

sudo docker pull mysql:latest

sudo docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql

sudo docker ps -a

sudo docker exec -it mysql-test bash

mysql -u root -p

CREATE DATABASE `po_go`

exit

exit

sudo docker pull redis

sudo docker run --name redis-lab -p 6379:6379 -d redis

sudo docker exec -it redis-lab bash

redis-cli

Ping

exit

exit
