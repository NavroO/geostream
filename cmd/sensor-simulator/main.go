package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/google/uuid"
)

type SensorData struct {
	ID          string    `json:"id"`
	SensorID    string    `json:"sensor_id"`
	Timestamp   time.Time `json:"timestamp"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Depth       float64   `json:"depth"`
	Temperature float64   `json:"temperature"`
	Pressure    float64   `json:"pressure"`
	Vibration   float64   `json:"vibration"`
	Moisture    float64   `json:"moisture"`
	DataType    string    `json:"data_type"`
	Unit        string    `json:"unit"`
	Value       float64   `json:"value"`
}

func main() {
	var wg sync.WaitGroup
	slog.Info("App run")
	wg.Go(func() {
		for {
			sensor := &SensorData{
				ID:          uuid.NewString(),
				SensorID:    uuid.NewString(),
				Timestamp:   time.Now(),
				Latitude:    randFloat(-90, 90),
				Longitude:   randFloat(-180, 180),
				Depth:       randFloat(0, 5000),
				Temperature: randFloat(-20, 120),
				Pressure:    randFloat(90000, 110000),
				Vibration:   randFloat(0, 10),
				Moisture:    randFloat(0, 100),
				DataType:    "geothermal",
				Unit:        "C",
				Value:       randFloat(0, 100),
			}

			jsonData, _ := json.MarshalIndent(sensor, "", "  ")
			fmt.Println(string(jsonData))

			time.Sleep(2 * time.Second)
		}
	})

	wg.Wait()
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
