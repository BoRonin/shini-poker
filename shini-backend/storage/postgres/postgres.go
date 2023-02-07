package postgres

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DB     *pgxpool.Pool
	counts int64
)

const DBTimeout = time.Second * 3

func openDB(dsn string) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return db, nil
}
func ConnectToDB() {
	dsn := os.Getenv("DSN")
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgress not yet ready...")
			counts++
		} else {
			log.Println("Connected to Postgres")
			SetupDbTable(connection)
			DB = connection
			return
		}
		if counts > 10 {
			log.Println(err)
			return
		}
		log.Println("Waiting for 2 seconds...")
		time.Sleep(time.Second * 2)
		continue
	}
}
func SetupDbTable(conn *pgxpool.Pool) {
	game := `SELECT EXISTS
    (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'game');`
	var result bool
	_ = conn.QueryRow(context.Background(), game).Scan(&result)

	if !result {
		c, ioErr := os.ReadFile("storage/getUp.sql")
		if ioErr != nil {
			log.Println(ioErr)
		}

		sql := string(c)
		_, err := conn.Exec(context.Background(), sql)
		if err != nil {
			fmt.Printf("Couldn't run query:%s", err.Error())
			return
		}

	}

}
