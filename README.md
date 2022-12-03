# Documentation
This project is a back-end application for personal website, it is recommended to be used together with the [front-end application](https://github.com/jdfcitecky/po_react). This project provides the basic mainstream functions of most websites, whether it is the basic database addition, deletion, modification and query system, membership system, display page, introduction page, comment system, background management, data chart, one-click restore, back-end file server , backend caching and even live chat rooms are fully implemented.

The website runs at http://www.chuanhsincho.com/ on AWS．
-----------------------------------------------------
# Below is the commands that show how to run this app at linux

Get docker.
#### `sudo apt-get install docker.io`

### Install mysql in docker
Get mysql image.
#### `sudo docker pull mysql:latest`
Run mysql image and assign the password.
#### `sudo docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql`
Check if mysql up.
#### `sudo docker ps -a`
Get into the mysql terminal.
#### `sudo docker exec -it mysql-test bash`
Login nysql.
#### `mysql -u root -p`
Create the database for backend service.
#### `CREATE DATABASE \`po_go\``
Exit the mysql command line.
#### `exit`
Exit the container terminal.
#### `exit`

### Install redis in docker
Get redis image.
#### `sudo docker pull redis`
Run redis image.
#### `sudo docker run --name redis-lab -p 6379:6379 -d redis`
Get into redis container terminal.
#### `sudo docker exec -it redis-lab bash`
Get into redis command line.
#### `redis-cli`
Check if redis is working.
#### `Ping`
Exit redis command line.
#### `exit`
Exit container terminal.
#### `exit`

### Install git
Get git.
#### `sudo apt-get install git-all`


### Get backend app
Clone the repository. You can create a directory for this projects first using `mkdir` and `cd` go into the directory.
#### `git clone https://github.com/jdfcitecky/po_go.git`
Rename the conf file.
#### `mv conf_template.yaml conf.yaml`
Modify the configuration. You can use any editor to do this step.
#### `vim conf.yaml`
+ Check the part of db, let host and password is correct in your machine.
+ The IP dail in docker should be 172.17.0.1
+ Change the password of mysql root.
Build docker image.
#### `sudo docker build . -t po_go`
Run the backend image.
#### `sudo docker run -p 4000:4000 -i -t -d po_go`

### Get frontend app
Clone the repository. You can create a directory for this projects first using `mkdir` and `cd` go into the directory.
#### `git clone https://github.com/jdfcitecky/portfolio_react_2022.git`
Modify the configuration. You can use any editor to do this step.
#### `vim .env`
+ Ensure REACT_APP_API_ADDRESS point to where your backend app run, the backend may run on `localhost:4000`.
+ Change the backend address to elastic ip, if you deploy on AWS.
Build docker image.
#### `sudo docker build . -t po_react`
Run the frontend image.
#### `sudo docker run -p 3000:3000 -i -t po_react`
+ Since some machine such as AWS do not have the server that react app use, if add `-d` flag, the app will hang at asking you to install the server. When it ask you to install, just type `y` and restart another terminal.

# Create an admin
Use GUI to do this, your forntend app may run at `localhost:3000`.

# Set database
Get into the mysql container terminal.
#### `sudo docker exec -it mysql-test bash`
Get into mysql command line.
#### `mysql -u root -p`
Use database.
#### `use po_go`
Let the member whoes id is 1 to be the admin, you can assign other member to be manager. But since the backend design, only the id=1 member can chat to everyone and everyone can only chat with the id=1 member.
#### `update members set is_manager=1 where id=1;`
Delete this record, or the admin cannot use chat room.
#### `delete from chat_room_aliases where id=2;`
-----------------------------------------------------
# About AWS setting

+ Get docker follow https://docs.docker.com/engine/install/ubuntu/
+ For nginx can follow https://ithelp.ithome.com.tw/articles/10221704
+ T2 micro will hang when build image of frontend app since the lack of memory
-----------------------------------------------------
# Known issues
+ Since this project was only developed for the desktop webpage at the beginning, even though it has been adapted for smaller screens, it still cannot display correctly on the mobile or a rather narrow window.