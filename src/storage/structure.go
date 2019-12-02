package storage

import "time"

type Event struct {
	Date        time.Time `json:"date"`
	DeviceId    string    `json:"deviceId"`
	Co2         int       `json:"cO2"`
	Temperature float32   `json:"temperature"`
	Humidity    int       `json:"humidity"`
}
