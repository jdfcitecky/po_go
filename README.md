# Getting Started with Create Go App
 Have to set db first


# below record how to run this app at linux

sudo apt-get install docker.io

### install mysql in docker
sudo docker pull mysql:latest

sudo docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql

sudo docker ps -a

sudo docker exec -it mysql-test bash

mysql -u root -p

CREATE DATABASE `po_go`

exit

exit

### install redis in docker
sudo docker pull redis

sudo docker run --name redis-lab -p 6379:6379 -d redis

sudo docker exec -it redis-lab bash

redis-cli

Ping

exit

exit
### install git
sudo apt-get install git-all

### get frontend app
git clone https://github.com/jdfcitecky/portfolio_react_2022.git

vim .env
+ change the backend address to elastic ip
sudo docker build . -t po_react
sudo docker run -p 3000:3000 -d po_react

### get backend app
git clone https://github.com/jdfcitecky/po_go.git

mv conf_template.yaml conf.yaml
vim conf.yaml
+ the ip dail in linux should be 172.17.0.1
+ change the password of mysql root
sudo docker build . -t po_go
sudo docker run -p 4000:4000 -i -t -d po_go

# For restart container
sudo docker container run ${container_name}

# Create an admin
use GUI to do this

# Set database
sudo docker exec -it mysql-test bash

mysql -u root -p

use po_go

update members set is_manager=1 where id=1;
delete from chat_room_aliases where id=2;

# about AWS setting

+ get docker follow https://docs.docker.com/engine/install/ubuntu/
+ t2 micro will hang at npm ci 