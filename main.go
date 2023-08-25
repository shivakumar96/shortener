package main

import (
	"sync"

	"url-shortner.com/backend/counter"
	"url-shortner.com/backend/worker"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go func() {
		counter.StartCounterServer()
		wg.Done()
	}()
	worker.WorkerEntry()
	wg.Wait()
}
