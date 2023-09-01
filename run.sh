#!/bin/sh

run_gateway(){
    sh ./scripts/start-api-gateway.sh 
}

run_counter(){
    sh ./scripts/start-counter.sh 
}

run_worker(){
    sh ./scripts/start-worker.sh 
}

run_all(){
    sh ./scripts/start-all-services.sh
}


run_db_scripts(){
    case $1 in 
        "setup") ./scripts/docker-db-setup.sh $1
        ;;
        "up" | "down" | "restart") ./scripts/db.sh $1
        ;;
        *) print_help
        ;;
    esac

}

print_help(){
    echo "-----------------------------------------------------------------------"
    echo "Follow the instruction to run the url server"
    echo "-----------------------------------------------------------------------"
    echo "./run.sh --<arg> <val>"
    echo "arg : urlservef or db. eg ./run.sh --db  <val> or ./run.sh --urlserver"
    echo "-----------------------------------------------------------------------"
    echo "--db <val> cann be setup, up, down, restart"
    echo "example ./run.sh --db setup or ./run.sh --db up or ./run.sh --db down"
    echo "or ./run.sh --db restart"
    echo " 1) setup : will install docker and docker compose"
    echo " 2) up: will run the Mysql docker conatiner"
    echo " 3) down: will stop the Mysql docker conatiner (Note: data will be loast)"
    echo " 4) restart: will restart the Mysql docker conatiner (Note: data will be loast)"
    echo "-----------------------------------------------------------------------"
    echo "--urlserver <val> cann be setup, gateway, counter, worker"
    echo "example ./run.sh --urlserver gateway or ./run.sh --urlserver counter"
    echo "or ./run.sh --urlserver worker"
    echo " 1) gateway : will start the API gateway server (microservice)"
    echo " 2) counter: will start the counter server (microservice)"
    echo " 3) worker: will start the worker server (microservice)"
    echo " 4) all: will start all three services, gateway, counter, worker"
    echo "-----------------------------------------------------------------------"
    echo "Note: configuration for the server can be modified in config.yml file"
    echo "-----------------------------------------------------------------------"
}

case $1 in 
    --urlserver) 
        case $2 in
            "gateway") run_gateway
            ;;
            "counter") run_counter
            ;;
            "worker") run_worker
            ;;
            "all")  run_all
            ;;
            *) print_help
            ;;
        esac
    ;;
    --db) run_db_scripts $2
    ;;
    *) print_help
    ;;
esac