package postgres

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// const (
// 	host     =
// 	port     = "5432"
// 	user     = "postgres"
// 	password = "abc123"
// 	dbname   = "postgres"
// )

var pool *pgxpool.Pool

func InitDbPool() {

	// get the database connection URL.
	// usually, this is taken as an environment variable as in below commented out code
	// databaseUrl = os.Getenv("DATABASE_URL")

	// for the time being, let's hard code it as follows.
	// ensure to change values as needed.
	databaseUrl := "host=" + DBConfig.Host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname

	// this returns connection pool
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Fprintf(os.Stderr, "Connection Successfull .")

	}

	// to close DB pool
	defer dbPool.Close()

}

func GetPool() *pgxpool.Pool {
	if pool == nil {
		InitDbPool()
	}

	connectedPoolSize := pool.AcquireAllIdle(context.Background())
	for connectedPoolSize == nil {
		// log.Println("Pg Connection Lost")
		pool.Close()
		time.Sleep(2 * time.Second)
		// log.Print("Reconnecting...")
		InitDbPool()
		connectedPoolSize = pool.AcquireAllIdle(context.Background())
	}
	return pool
}
