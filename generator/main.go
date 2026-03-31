package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
)

type Telemetry struct {
	Latency   int       `json:"latency"`
	CPU       int       `json:"cpu"`
	Players   int       `json:"players"`
	Timestamp time.Time `json:"timestamp"`
}

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	for {
		data := Telemetry{
			Latency:   rand.Intn(100),
			CPU:       rand.Intn(100),
			Players:   rand.Intn(5000),
			Timestamp: time.Now(),
		}

		jsonData, _ := json.Marshal(data)

		rdb.LPush(ctx, "telemetry", jsonData)

		time.Sleep(1 * time.Second)
	}
}
