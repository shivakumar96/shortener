#!/bin/sh

macOS_setup() {
    if [ -f /usr/local/bin/docker ]
    then
        return 0    
    fi
    brew install docker
    brew install docker-compose
}

linux_setup() {
    if [ -f /usr/local/bin/docker ]
    then
        return 0    
    fi
    sudo apt-get update
    sudo apt-get install docker-ce
    sudo apt-get install docker-compose
}

case "$OSTYPE" in
  darwin*)  macOS_setup ;; 
  linux*)   linux_setup ;;
  *)        echo "OS NOT SUPPOSRTED" 
            exit 1;;
esac

docker-compose -f db-configuration/mysql-docker-compose.yaml  up -d
