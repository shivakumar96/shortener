package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"url-shortner.com/backend/db"
	"url-shortner.com/backend/utils"
	"url-shortner.com/backend/worker"
)

var config *utils.Config

// randomly select a worker node
func selectRandomWorker() string {
	randnum := rand.Intn(config.Worker.Count)
	var port, _ = strconv.Atoi(config.Worker.Port)
	port += randnum
	workerURL := "http://" + config.Worker.Host + ":" + strconv.Itoa(port) + "/tinyurl"
	return workerURL
}

func getFullURL(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	shortURL := mux.Vars(request)["shortURL"]
	workerURL := selectRandomWorker()
	workerURL = workerURL + "/" + shortURL
	workerRequest, _ := http.NewRequest(http.MethodGet, workerURL, nil)
	var client = &http.Client{Timeout: time.Second * 40}
	resp, err := client.Do(workerRequest)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	if resp.StatusCode != http.StatusOK {
		writer.WriteHeader(resp.StatusCode)
		return
	}
	var urls db.Tiny2LongURL
	json.NewDecoder(resp.Body).Decode(&urls)
	json.NewEncoder(writer).Encode(&urls)

}

func createShortURL(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	workerURL := selectRandomWorker()

	var longURLReqBody worker.LongURLStruct
	json.NewDecoder(request.Body).Decode(&longURLReqBody)

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&longURLReqBody)

	log.Println(buf)

	workerRequest, _ := http.NewRequest(http.MethodPost, workerURL, &buf)
	var client = &http.Client{Timeout: time.Second * 40}
	resp, err := client.Do(workerRequest)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	if resp.StatusCode != http.StatusOK {
		writer.WriteHeader(resp.StatusCode)
		return
	}
	var urls db.Tiny2LongURL
	json.NewDecoder(resp.Body).Decode(&urls)
	json.NewEncoder(writer).Encode(&urls)
}

func StartAPIGateway() {
	config, _ = utils.ReadConfig()

	router := mux.NewRouter()
	router.HandleFunc("/{shortURL}", getFullURL).Methods("GET")
	router.HandleFunc("/", createShortURL).Methods("POST")

	config, _ := utils.ReadConfig()
	port, _ := strconv.Atoi(config.Server.Port)

	serverURL := config.Worker.Host + ":" + strconv.Itoa(port)

	srv := &http.Server{
		Addr:         serverURL,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 30,
		Handler:      router,
	}
	fmt.Printf("API Gateway server is listening at %v \n", serverURL)
	log.Fatal(srv.ListenAndServe())
}
