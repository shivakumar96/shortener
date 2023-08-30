package counter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"url-shortner.com/backend/utils"
)

var workerCount int
var maxWorkerCount int
var rangeValue int

var mutex sync.Mutex

type WorkedId struct {
	WorkedID int `json:"workerID"`
}

type Range struct {
	StartRange int `json:"startRange"`
	EndRange   int `json:"endRange"`
}

// function to join the worker cluster
func joinMethod(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// make ihe counter safe
	mutex.Lock()

	if workerCount >= maxWorkerCount {
		log.Println(workerCount, maxWorkerCount)
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Internal server error: worker exceeded")
		return
	}

	val := workerCount
	workerCount++
	mutex.Unlock()

	var worker = WorkedId{WorkedID: val}
	log.Println("worked join with ID", worker)
	err := json.NewEncoder(writer).Encode(worker)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

// function to return the range
func getRange(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["workedId"])

	if id >= workerCount || err != nil {
		log.Println(workerCount, maxWorkerCount)
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Internal server error: worker does not exists")
		return
	}

	startVal := rangeValue * id
	endVal := rangeValue * (id + 1)

	var ran = Range{StartRange: startVal, EndRange: endVal}

	err = json.NewEncoder(writer).Encode(ran)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
	}

}

// entry method for counter
func StartCounterServer() {

	var config *utils.Config
	config, err := utils.ReadConfig()

	if err != nil {
		fmt.Print(err)
		fmt.Print("Error loading config")
	}

	workerCount = 0
	maxWorkerCount = config.Counter.MaxCount
	rangeValue = config.Counter.Ranges
	conf, _ := json.Marshal(config.Counter)
	log.Printf("counter configuration: %v", string(conf))

	router := mux.NewRouter()
	router.HandleFunc("/join", joinMethod).Methods("GET")
	router.HandleFunc("/range/{workedId}", getRange).Methods("GET")

	counterURL := config.Counter.Host + ":" + config.Counter.Port

	srv := &http.Server{
		Addr:         counterURL,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 30,
		Handler:      router,
	}
	fmt.Printf("Counter listening at %v \n", counterURL)
	log.Fatal(srv.ListenAndServe())

}
