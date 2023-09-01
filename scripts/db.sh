#!/bin/sh

case $1 in
    "up") docker-compose -f db-configuration/mysql-docker-compose.yaml  up -d 
    ;;
    "down") docker-compose -f db-configuration/mysql-docker-compose.yaml  down 
    ;;
    "restart") docker-compose -f db-configuration/mysql-docker-compose.yaml  restart -d 
    ;;
    *)
    echo "command not supported"
    echo "try db.sh <args (up or down or restart)>"
    ;;
esac


