# **URL-Shortener**

## Golang based Microservice application project:
 A Full-Stack project developed using API REST-based microservices architecture created using Golang, ReactJS, and MySQL.

## Requirements:
- Golang  [here](https://go.dev/doc/install)
- Modules/libraries :
    - Gorilla Mux 
    - Gorm
    - Mysql driver

- ReactJS
- MysQL Database (Docker container)

### Note: This Project is an MVP, and there are scopes to improve this Project. 
### cuttent version is 1.0

## Architecture

<img src="https://github.com/shivakumar96/url-shortener/blob/main/architecture/tinyURL_architecture.png" width="800" height="700">

# How to set up ?

## Clone Repository from GitHub
Run the below command to clone code from github
```
git clone https://github.com/shivakumar96/url-shortener.git
```

## build the code (Optional)
Executable file 'backend' is present in the executables directory, of not run the below command.
```
go build -o ./executables
```

## Install docker and docker-compose (Optional)
This project will use the Mysql docker image for db setup. <br>
Run the below command from the url-shortner directory
```
./run.sh --db setup
```

## Run MySQL DB container
To run the MySQL DB, run the below command from the url-shortner directory
```
./run.sh --db up
```

To stop the MySQL DB, run the below command from the url-shortner directory
```
./run.sh --db down
```

## Run Micro services
To run the API-Gateway microservice, run the below command from the url-shortner directory
```
./run.sh --urlserver gateway
```

To run the Counter microservice, run the below command from the url-shortner directory
```
./run.sh --urlserver counter
```

To run the Worker microservice, run the below command from the url-shortner directory
```
./run.sh --urlserver worker
``` 

To run all microservices, run the below command from the url-shortner directory
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