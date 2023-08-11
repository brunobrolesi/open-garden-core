package db

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
)

var timescaleSingleInstance *pgx.Conn
var timescaleOnce = &sync.Once{}

func GetTimescaleInstance() *pgx.Conn {
	timescaleOnce.Do(func() {
		conn, err := pgx.Connect(context.Background(), os.Getenv("TIMESCALE_URL"))
		if err != nil {
			log.Fatal("timescale db connection fail")
		}
		timescaleSingleInstance = conn
	})

	return timescaleSingleInstance
}
