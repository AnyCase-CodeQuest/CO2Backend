package main

import (
	"encoding/json"
	"fmt"
	"github.com/AnyCase-Company-LTD/CO2Backend/src/message"
	"github.com/AnyCase-Company-LTD/CO2Backend/src/static"
	"github.com/AnyCase-Company-LTD/CO2Backend/src/storage"
	"github.com/AnyCase-Company-LTD/CO2Backend/src/values"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

import log "github.com/sirupsen/logrus"

type myNotFoundHandler struct {
}

func (*myNotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "not found page: %s", r.URL.Path)
	log.Warnf("404 page: %s; IP: %s; body: %s", r.URL.Path, r.RemoteAddr, r.Body)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newEvent storage.Event
	newEvent.Date = time.Now()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	err = storage.Create(newEvent)

	if os.Getenv(values.EnvSendToQueue) == "1" {
		msg, _ := json.Marshal(newEvent)
		go message.SendMessageToQueue(msg)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	json.NewEncoder(w).Encode(newEvent)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
}

func main() {
	setupEnv()
	router := mux.NewRouter().StrictSlash(true)
	spa := static.SpaHandler{StaticPath: "build", IndexPath: "index.html"}
	router.PathPrefix("/spa").Handler(spa)
	router.HandleFunc("/", homeLink).Methods(http.MethodGet)
	router.HandleFunc("/event", createEvent).Methods(http.MethodPost).Headers("Content-Type", "application/json")
	router.HandleFunc("/event/latest", getLatestEvent).Methods(http.MethodGet).Headers("Content-Type", "application/json")
	router.HandleFunc("/api/v1/sensor/{id}/latest", getLatestSensor).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/sensor-list", getSensorList).Methods(http.MethodGet, http.MethodOptions)
	router.NotFoundHandler = &myNotFoundHandler{}
	log.WithError(http.ListenAndServe(":8080", router)).Fatal()
}

func getLatestEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	singleEvent, err := storage.GetLatest()
	if err != nil {
		log.Error("can't get latest value", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(singleEvent)
	}
}

func getLatestSensor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	setupCORS(w)

	if r.Method == http.MethodGet {
		singleEvent, err := storage.GetLatestBy(vars["id"])
		if err != nil {
			log.Error("can't get latest value", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getSensorList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	setupCORS(w)

	if r.Method == http.MethodGet {
		sensorList := storage.GetSensorList()

		json.NewEncoder(w).Encode(sensorList.Data)
	}

}

func setupCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Method", "GET")
	w.Header().Set("Accept", "*/*")
}

func setupEnv() {
	switch environment := os.Getenv(values.EnvEnv); environment {
	case values.EnvProd:
		log.SetLevel(log.WarnLevel)
		break
	case values.EnvDev:
		log.SetLevel(log.DebugLevel)
	default:
		log.Warnf("No env variable set: \"%s\" Set as \"dev\" default", values.EnvEnv)
		err := os.Setenv(values.EnvEnv, "dev")
		if err != nil {
			log.Error(err)
		}
		log.SetLevel(log.DebugLevel)
	}
}
