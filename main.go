package main

import (
	"fmt"
	"sync"

	"url-shortner.com/backend/db"
)

var wg sync.WaitGroup

func main() {
	/*	wg.Add(1)
		go func() {
			counter.StartCounterServer()
			wg.Done()
		}()
		worker.WorkerEntry()
		wg.Wait()
	*/
	db.ConnectToDB()
	url := db.Tiny2LongURL{Tinyurl: "abc", Longurl: "longurl"}
	db.AddURL(&url)

	fmt.Println(db.GetFullURL("abc"))
}
