package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
	_ "github.com/lib/pq"
)

type Telemetry struct {
	Latency   int
	CPU       int
	Players   int
	Timestamp string
}

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	connStr := "user=telemetry password=telemetry dbname=telemetry sslmode=disable host=postgres"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		result, err := rdb.BRPop(ctx, 0, "telemetry").Result()
		if err != nil {
			continue
		}

		var t Telemetry
		json.Unmarshal([]byte(result[1]), &t)

		_, err = db.Exec(
			"INSERT INTO telemetry(latency, cpu, players, timestamp) VALUES ($1,$2,$3,$4)",
			t.Latency, t.CPU, t.Players, t.Timestamp,
		)

		if err != nil {
			log.Println(err)
		}
	}
}
