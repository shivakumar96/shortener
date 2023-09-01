package main

import (
	"log"
	"sync"
	"time"

	gateway "url-shortner.com/backend/API-Gateway"
	"url-shortner.com/backend/counter"
	"url-shortner.com/backend/utils"
	"url-shortner.com/backend/worker"
)

var wg sync.WaitGroup
var config *utils.Config

func main() {
	config, _ = utils.ReadConfig()
	wg.Add(1 + config.Worker.Count)
	go func() {
		counter.StartCounterServer()
		wg.Done()
	}()

	for i := 0; i < config.Worker.Count; i++ {
		log.Println("staring worker")
		time.Sleep(time.Second * 5)

		go func() {
			worker.WorkerEntry()
			wg.Done()
		}()
	}

	log.Println("staring gateway")
	gateway.StartAPIGateway()
	wg.Wait()

	//db.ConnectToDB()
	//url := db.Tiny2LongURL{Tinyurl: "abc", Longurl: "longurl"}
	//db.AddURLIfAbsent(&url)

	//fmt.Println(db.GetFullURL("abc"))
}
