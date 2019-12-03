package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"storage"

	"github.com/gorilla/mux"
)

type myNotFoundHandler struct {
}

func (*myNotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "not found page: %s", r.URL.Path)
	log.Println(fmt.Sprintf("404 page: %s; IP: %s; body: %s", r.URL.Path, r.RemoteAddr, r.Body))
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

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	json.NewEncoder(w).Encode(newEvent)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	//eventID := mux.Vars(r)["id"]

	//for _, singleEvent := range events {
	//	if singleEvent.DeviceId == eventID {
	//		json.NewEncoder(w).Encode(singleEvent)
	//	}
	//}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink).Methods(http.MethodGet)
	router.HandleFunc("/event", createEvent).Methods(http.MethodPost).Headers("Content-Type", "application/json")
	router.HandleFunc("/event/latest", getLatestEvent).Methods(http.MethodGet).Headers("Content-Type", "application/json")
	router.HandleFunc("/events/{id}", getOneEvent).Methods(http.MethodGet).Headers("Content-Type", "application/json")
	router.NotFoundHandler = &myNotFoundHandler{}
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getLatestEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	singleEvent, err := storage.GetLatest()
	if err != nil {
		log.Println("can't get latest value", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(singleEvent)
	}

}
