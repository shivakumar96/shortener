# **URL-Shortener**

## Golang based Microservice application:
 A Full-Stack project developed using API REST-based microservices architecture created using Golang, ReactJS, and MySQL.

## Requirements:
- Golang, Install it from [here](https://go.dev/doc/install)
- Modules/libraries :
    - Gorilla Mux 
    - Gorm
    - Mysql driver

- ReactJS
- MysQL Database (Docker container)

### Note: This Project is an MVP, and there are scopes to improve this Project. 
### This repository contains the backend code the URL-shortener project. The current version of the backend is v1.0.
### url-shortener-ui: UI for URL-shortener Project
The frontend code for the URL-shortener Project can be found in the repository mentioned in the below link  <br>
[https://github.com/shivakumar96/url-shortener-ui](https://github.com/shivakumar96/url-shortener-ui)
 

## Architecture

<img src="https://github.com/shivakumar96/url-shortener/blob/main/architecture/tinyURL_architecture.png" width="800" height="700">

# How to set up ?

## Clone Repository from GitHub
Run the below command to clone code from GitHub
```
git clone https://github.com/shivakumar96/url-shortener.git
```

## Download Modules 
Downlod go modules/libraries
```
go mod download
```

## build the code (Optional)
The executable file 'backend' is in the executables directory. If not, run the below command.
```
go build -o ./executables
```

## Install docker and docker-compose (Optional)
This project will use the Mysql docker image for the DB setup. <br>
Run the below command from the url-shortener directory
```
./run.sh --db setup
```

## Run MySQL DB container
To run the MySQL DB, run the below command from the url-shortener directory
```
./run.sh --db up
```

To stop the MySQL DB, run the below command from the url-shortener directory
```
./run.sh --db down
```

## Modify Configuration (Optional)
Update the configuration file ```config.yml``` <br>
If DB configurations are modified, the same values should be reflected in ``` db-configuration/mysql-docker-compose.yaml```

## Run Micro services
To run the API-Gateway microservice, run the below command from the url-shortener directory
```
./run.sh --urlserver gateway
```

To run the Counter microservice, run the below command from the url-shortener directory
```
./run.sh --urlserver counter
```

To run the Worker microservice, run the below command from the url-shortener directory
```
./run.sh --urlserver worker
``` 

To run all microservices, run the below command from the url-shortener directory
```
./run.sh --urlserver all
``` 

## Need help?
Run the below command from the url-shortner directory, to get help on how to execute the service.
```
./run.sh --help
``` 

## output
<img src="https://github.com/shivakumar96/url-shortener/blob/main/architecture/output.png" width="600" height="400">