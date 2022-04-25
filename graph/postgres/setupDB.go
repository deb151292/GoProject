package postgres

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	dbconst "github.com/deb151292/gqlgen-todos/graph/DBconst"
	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

func InitDbPool() {

	// get the database connection URL.
	// usually, this is taken as an environment variable as in below commented out code
	// databaseUrl = os.Getenv("DATABASE_URL")

	// for the time being, let's hard code it as follows.
	// ensure to change values as needed.

	databaseUrl := "host=" + dbconst.Host + " port=" + strconv.Itoa(dbconst.Port) + " user=" + dbconst.User + " password=" + dbconst.Password + " dbname=" + dbconst.Dbname

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
