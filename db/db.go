package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/lib/pq"
)

func Open() *sql.DB {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		url = "postgres://enxzlztruhxwlg:bf1f4ccc6c1e2393695541093e70c2c17b994deda7163463734994f232c178e8@ec2-54-225-88-191.compute-1.amazonaws.com:5432/d31ibrodp1403o?sslmode=require"
	}
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Connected!")
	}

	return db
}
