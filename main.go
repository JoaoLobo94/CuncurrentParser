package main

import (
	"database/sql"
	"fmt"
	"github.com/JoaoLobo94/donut_test/util"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgresql://root:1234abcd@localhost:5433/donut_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	fmt.Println("How many random transactions would you like to generate for this test?: ")
	fmt.Println(util.RandTransactionAmounts())
}
