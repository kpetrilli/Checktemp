package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yryz/ds18b20"
)

// What the user gets
type UserResponse struct {
	Measurements []SensorData `json:"data"`
}

type SensorData struct {
	SensorID    int     `json:"sensor"`
	Temperature float64 `json:"temperature"`
}

// getTemperature gets the temperature of all registered sensors
func getTemperature() []SensorData {
	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	var payload []SensorData
	for id, sensor := range sensors {
		temp, err := ds18b20.Temperature(sensor)
		if err != nil {
			panic(err)
		}
		results := SensorData{
			SensorID:    id,
			Temperature: temp,
		}
		payload = append(payload, results)
	}
	return payload

}

// Log source address and respond
func APIHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("Request from:", r.RemoteAddr)

	response := UserResponse{
		Measurements: getTemperature(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Barebones HTTP serving
func main() {
	http.HandleFunc("/", APIHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}
