package db

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
)

var lock = &sync.Mutex{}

var singleInstance *pgx.Conn

func GetInstance() *pgx.Conn {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
			if err != nil {
				log.Fatal("database connection fail")
			}
			singleInstance = conn
		}
	}
	return singleInstance
}
