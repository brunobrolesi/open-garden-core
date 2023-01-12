package db

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
)

var singleInstance *pgx.Conn
var once = &sync.Once{}

func GetInstance() *pgx.Conn {
	once.Do(func() {
		conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatal("database connection fail")
		}
		singleInstance = conn
	})

	return singleInstance
}
