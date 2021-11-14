package util

import (
	"context"
	"fmt"
	"github.com/JoaoLobo94/donut_test/db/sqlc"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"time"
)

func Generate(ctx context.Context, queries *db.Queries) {
	initializePromptWithData(ctx, queries, generateUser(ctx, queries))
}

func initializePromptWithData(ctx context.Context, queries *db.Queries, user db.User) {
	rand.Seed(time.Now().UnixNano())
	var numberToGenerate int
	fmt.Scanln(&numberToGenerate)
	min := 0.0
	max := 99.9
	result := make([]float64, numberToGenerate)
	fmt.Println("Seeding... Please be patient")
	for i := range result {
		result[i] = min + rand.Float64()*(max-min)
		queries.CreateAction(ctx, db.CreateActionParams{
			Amount: result[i],
			UserID: user.ID,
		})
	}
	
	fmt.Println("You just created ",len(result)," fake bank transactions")
	startBatches(ctx, queries, user.ID)
	Prompt(ctx, queries, user.ID)
}

func generateUser(ctx context.Context, queries *db.Queries) db.User {
	fmt.Printf("Please tell me what is your name -->  ")
	var nameOfUser string
	fmt.Scanln(&nameOfUser)
	user, err := queries.CreateUser(ctx, nameOfUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hey " + nameOfUser + " lets get started. How many bank transactions would you to seed? --> ")
	return user
}

func startBatches(ctx context.Context, queries *db.Queries, current_user int32) {
	queries.CreateBatch(ctx, db.CreateBatchParams{UserID: current_user})
}
