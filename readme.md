## Introduction

Here's **what** you should do to run this project
For Linux/MacOs:

- install docker, docker-compose
- run `make rebuild`
- run docker ps -a
- copy the id of the first container with name "xxx-app-1"
- run this command `docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' first-id`
- copy the ip address
- add the your /etc/hosts with value of the hosts from .env
- go to /sql/migrate
- copy and execute the whole sql command
- rerun `make rebuild`

For Windows:

- follow this instruction https://linuxhint.com/run-makefile-windows/
- install docker, docker-compose
- run `make rebuild`
- run docker ps -a
- copy the id of the first container with name "xxx-app-1"
- run this command `docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' first-id`
- copy the ip address
- add the your `C:\Windows\System32\drivers\etc\hosts` with value of the hosts from .env
- go to /sql/migrate
- copy and execute the whole sql command
- rerun `make rebuild`
