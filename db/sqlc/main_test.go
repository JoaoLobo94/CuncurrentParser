package db

import (
	"os"
	"log"
	"database/sql"
	"testing"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var testQueries *Queries

func init() {
	err := godotenv.Load("../../.env")
    
	if err != nil {
	    log.Fatal("Error loading .env file")
	}
    }
func TestMain(m *testing.M) {
	var user, password, sslmode = os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("SSLMODE")
	conn, err := sql.Open("postgres", "postgresql://" + user + ":" + password + "@localhost:5433/donut_db_test?sslmode=" + sslmode)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}
