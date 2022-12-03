# Documentation
This project is a back-end application for personal display webpage, it is recommended to be used together with the [front-end application](https://github.com/jdfcitecky/po_react). This project provides the basic mainstream functions of most websites, whether it is the basic database addition, deletion, modification and query system, membership system, display page, introduction page, comment system, background management, data chart, one-click restore, back-end file server , backend caching and even live chat rooms are fully implemented.

-----------------------------------------------------
# Below record how to run this app at linux

Get docker.
#### `sudo apt-get install docker.io`

### install mysql in docker
Get mysql image.
#### `sudo docker pull mysql:latest`

#### `sudo docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql`

#### `sudo docker ps -a`

#### `sudo docker exec -it mysql-test bash`

#### `mysql -u root -p`

#### `CREATE DATABASE \`po_go\``

#### `exit`

#### `exit`

### install redis in docker
#### `sudo docker pull redis`

#### `sudo docker run --name redis-lab -p 6379:6379 -d redis`

#### `sudo docker exec -it redis-lab bash`

#### `redis-cli`

#### `Ping`

#### `exit`

#### `exit`

### install git

#### `sudo apt-get install git-all`

### get frontend app

#### `git clone https://github.com/jdfcitecky/portfolio_react_2022.git`

#### `vim .env`
+ change the backend address to elastic ip
#### `sudo docker build . -t po_react`

#### `sudo docker run -p 3000:3000 -i -t po_react`

### get backend app
#### `git clone https://github.com/jdfcitecky/po_go.git`

#### `mv conf_template.yaml conf.yaml`
#### `vim conf.yaml`
+ the ip dail in linux should be 172.17.0.1
+ change the password of mysql root
#### `sudo docker build . -t po_go`
#### `sudo docker run -p 4000:4000 -i -t -d po_go`


# Create an admin
use GUI to do this

# Set database
#### `sudo docker exec -it mysql-test bash`

#### `mysql -u root -p`

#### `use po_go`

#### `update members set is_manager=1 where id=1;`
#### `delete from chat_room_aliases where id=2;`
-----------------------------------------------------
# About AWS setting

+ Get docker follow https://docs.docker.com/engine/install/ubuntu/
+ For nginx can follow https://ithelp.ithome.com.tw/articles/10221704
+ T2 micro will hang when build image of frontend app since the lack of memory
-----------------------------------------------------
# Known issues
+ Since this project was only developed for the desktop webpage at the beginning, even though it has been adapted for smaller screens, it still cannot display correctly on the mobile or a rather narrow window.