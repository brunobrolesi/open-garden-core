package db

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
)

var postgresSingleInstance *pgx.Conn
var postgresOnce = &sync.Once{}

func GetPostreSQLInstance() *pgx.Conn {
	postgresOnce.Do(func() {
		conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRESQL_URL"))
		if err != nil {
			log.Fatal("postgres db connection fail")
		}
		postgresSingleInstance = conn
	})

	return postgresSingleInstance
}
