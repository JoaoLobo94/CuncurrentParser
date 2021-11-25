package util

import (
	"context"
	"database/sql"
	"github.com/JoaoLobo94/donut_test/db/sqlc"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pioz/faker"
	"log"
	"os"
	"testing"
)

func TestBatchAllAmounts(t *testing.T) {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var user, password, sslmode = os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("SSLMODE")
	conn, err := sql.Open("postgres", "postgresql://"+user+":"+password+"@localhost:5433/donut_db_test?sslmode="+sslmode)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	queries := db.New(conn)
	name := faker.Username()
	appUser, _ := queries.CreateUser(context.Background(), name)
	actionArg := db.CreateActionParams{
		Amount: faker.Float64InRange(1.11, 1.12),
		UserID: appUser.ID,
	}

	for i := 0; i < 250; i++ {
		queries.CreateAction(context.Background(), actionArg)

	}
	batchArg := db.CreateBatchParams{
		Dispatched: false,
		Amount:     faker.Float64(),
		UserID:     appUser.ID,
	}

	batch, _ := queries.CreateBatch(context.Background(), batchArg)
	batchAllAmounts(context.Background(), queries, appUser.ID, batch.ID)
	allTransactions, _ := queries.ListTransactions(context.Background())
	updatedBatch, _ := queries.GetBatch(context.Background(), batch.ID)
	if updatedBatch.Amount == 0 {
		log.Fatal("error, batch not updated")
	}
	if updatedBatch.Amount >= 100 && len(allTransactions)< 1{
		log.Fatal("transaction not broadcasted")
	}
	cleanUp(context.Background(), queries)
}
