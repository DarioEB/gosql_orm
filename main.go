package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	var pool *sql.DB
	var dsnType string
	const driver = "postgres"
	id := flag.Int64("id", 1, "person id to find")

	switch driver {
	case "mysql":
		dsnType = "devuser:devpass@tcp(127.0.0.1:3306)/gosql"
	case "postgres":
		dsnType = "host=localhost port=5432 dbname=gopsql user=devuser password=devpass sslmode=disable"
	default:
		dsnType = ""
	}

	dsn := flag.String("dsn", dsnType, "connection data source name")

	flag.Parse()

	if len(*dsn) == 0 {
		log.Fatal("missing dsn flag")
	}

	if *id == 0 {
		log.Fatal("missing id person")
	}

	fmt.Println(*dsn)
	fmt.Println(*id)

	pool, err := sql.Open(driver, *dsn)
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}
	defer pool.Close()

	pool.SetConnMaxIdleTime(3)
	pool.SetMaxIdleConns(3)
	pool.SetConnMaxLifetime(0)

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)

	go func() {
		<-appSignal
		stop()
	}()

	ping(ctx, pool)
}

func ping(ctx context.Context, pool *sql.DB) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect database: %v", err)
	}

	log.Println("connected to database", pool.Driver())
}
