package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	DeviceId    string  `json:"deviceId"`
	Co2         int     `json:"cO2"`
	Temperature float32 `json:"temperature"`
	Humidity    int     `json:"humidity"`
}

type allEvents []event

var events = allEvents{
	{
		DeviceId:    "fa238a69-03ab-40d1-a51c-eb384844d243",
		Co2:         500,
		Temperature: 24.2,
		Humidity:    44,
	},
}

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
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.DeviceId == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink).Methods(http.MethodPost)
	router.HandleFunc("/event", createEvent).Methods(http.MethodPost).Headers("Content-Type", "application/json")
	router.HandleFunc("/events", getAllEvents).Methods(http.MethodGet).Headers("Content-Type", "application/json")
	router.HandleFunc("/events/{id}", getOneEvent).Methods(http.MethodGet).Headers("Content-Type", "application/json")
	router.NotFoundHandler = &myNotFoundHandler{}
	log.Fatal(http.ListenAndServe(":8080", router))
}
