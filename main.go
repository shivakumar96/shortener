package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	gateway "url-shortner.com/backend/API-Gateway"
	"url-shortner.com/backend/counter"
	"url-shortner.com/backend/utils"
	"url-shortner.com/backend/worker"
)

var wg sync.WaitGroup
var config *utils.Config

func runCounter() {
	counter.StartCounterServer()
	wg.Done()
}

func runWorker() {
	worker.WorkerEntry()
	wg.Done()
}

func runAPIGateway() {
	gateway.StartAPIGateway()
	wg.Done()
}

func helpDescription() {
	log.Println("command not suppoted")
	fmt.Println("run the program as ./backend <arg>")
	fmt.Println("arg = all, will runn all the services")
	fmt.Println("arg = gateway, will run the api gateway")
	fmt.Println("arg = counter, will run the counter")
	fmt.Println("arg = worker, will run the worker, (make sure counter is running before worker)")
}

func main() {
	log.Println("starting URL shortener...")
	config, _ = utils.ReadConfig()

	commands := os.Args[1:]

	log.Println(commands)
	if len(commands) == 0 {
		helpDescription()
		os.Exit(0)
	}

	switch commands[0] {
	case "all":
		wg.Add(2 + config.Worker.Count)
		go runCounter()
		for i := 0; i < config.Worker.Count; i++ {
			time.Sleep(time.Second * 5)
			go runWorker()
		}
		go runAPIGateway()

	case "counter":
		wg.Add(1)
		go runCounter()
	case "worker":
		wg.Add(1)
		go runWorker()
	case "gateway":
		wg.Add(1)
		go runAPIGateway()
	default:
		helpDescription()
	}
	wg.Wait()
}
