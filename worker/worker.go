package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"url-shortner.com/backend/counter"
	"url-shortner.com/backend/utils"
)

var mutex sync.Mutex
var id int
var rangeStruct counter.Range
var currentCount int

type shortURLStruct struct {
	ShortURL string `json:"shortURL"`
}

type LongURLStruct struct {
	LongURL string `json:"longURL"`
}

func createShortURL(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	log.Println("create request")
	mutex.Lock()
	uniqueNum := currentCount
	currentCount++
	if currentCount >= rangeStruct.EndRange {
		go initilazeWorker()
	} // reinitialize the server
	mutex.Unlock()
	shortURl := utils.ConvertIntToB64(uniqueNum)
	log.Println(shortURl)

	var sURl = shortURLStruct{ShortURL: shortURl}

	err := json.NewEncoder(writer).Encode(&sURl)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
	}

}

func getFullURL(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	log.Println("get full url request")

}

func shutdown(message string) {
	log.Println(message)
	os.Exit(1)
}

func request(request *http.Request, client *http.Client, body interface{}) {
	resp, err := client.Do(request)

	if err != nil {
		shutdown("cannot join the counter")
	} else if resp.StatusCode != http.StatusOK {
		shutdown("cannot create a worker")
	}
	json.NewDecoder(resp.Body).Decode(body)
}

func initilazeWorker() {
	var config, err = utils.ReadConfig()
	if err != nil {
		shutdown("cannot initilize server")
	}
	counterURL := "http://" + config.Counter.Host + ":" + config.Counter.Port
	var client = &http.Client{Timeout: time.Second * 40}

	// join the cluter
	joinURL := counterURL + "/join"
	joinRequest, _ := http.NewRequest(http.MethodGet, joinURL, nil)
	var worker counter.WorkedId
	request(joinRequest, client, &worker)
	mutex.Lock()
	id = worker.WorkedID
	log.Println("worker joined", id)
	mutex.Unlock()

	// obtain the range
	rangeURL := counterURL + "/range/" + strconv.Itoa(id)
	rangeRequest, _ := http.NewRequest(http.MethodGet, rangeURL, nil)
	mutex.Lock()
	request(rangeRequest, client, &rangeStruct)
	log.Println(rangeStruct)
	mutex.Unlock()

	//set the current range // locaking is not necessary during initialization
	mutex.Lock()
	currentCount = rangeStruct.StartRange + 1
	mutex.Unlock()
	log.Println("worker initilaiztion completed")

}

func WorkerEntry() {
	initilazeWorker()

	router := mux.NewRouter()
	router.HandleFunc("/tinyurl/{val}", getFullURL).Methods("GET")
	router.HandleFunc("/tinyurl", createShortURL).Methods("Post")

	config, _ := utils.ReadConfig()
	port, _ := strconv.Atoi(config.Worker.Port)
	mutex.Lock()
	port += id
	mutex.Unlock()

	workerURL := config.Worker.Host + ":" + strconv.Itoa(port)

	srv := &http.Server{
		Addr:         workerURL,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 30,
		Handler:      router,
	}
	fmt.Printf("Worker id %v listening at %v \n", id, workerURL)
	log.Fatal(srv.ListenAndServe())
}
