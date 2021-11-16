package util

import (
	"context"
	"fmt"
	"log"
	"math"

	"github.com/AlecAivazis/survey/v2"
	db "github.com/JoaoLobo94/donut_test/db/sqlc"
	"github.com/davecgh/go-spew/spew"
)

func Prompt(ctx context.Context, queries *db.Queries, currentUser int32, currentBatch int32) {
	res := ""
	prompt := &survey.Select{
		Message: "<---------What would you like to do now?--------->\n",
		Options: []string{"Start Over --> This will erase all database",
			"Start Batching low volume bank transactions. This will run concurrently",
			"Which user am I?",
			"View all batches",
			"View all dispatched batches",
			"View all transactions sent to the bank",
			"View all seeded bank user bank transactions to round up"},
	}
	survey.AskOne(prompt, &res)

	switch res {
	case "Start Over --> This will erase all database":
		startOver(ctx, queries, currentUser, currentBatch)
	case "Start Batching low volume bank transactions. This will run concurrently":
		startBatch(ctx, queries, currentUser, currentBatch)
	case "Which user am I?":
		whoAmI(ctx, currentUser, queries, currentBatch)
	case "View all batches":
		listBatches(ctx, queries, currentUser, currentBatch)
	case "View all dispatched batches":
		dispatchedBatches(ctx, queries, currentUser, currentBatch)
	case "View all transactions sent to the bank":
		listTransactions(ctx, queries, currentUser, currentBatch)
	case "View all seeded bank user bank transactions to round up":
		listActions(ctx, queries, currentUser, currentBatch)
	default:
		cleanUp(ctx, queries)
	}
}

func cleanUp(ctx context.Context, queries *db.Queries) {
	queries.DeleteActions(ctx)
	queries.DeleteBatches(ctx)
	queries.DeleteTransactions(ctx)
	queries.DeleteUsers(ctx)
}
func startOver(ctx context.Context, queries *db.Queries, currentUser int32, currentBatch int32) {
	fmt.Printf("Erasing....\n")
	cleanUp(ctx, queries)
	fmt.Printf("Erased \n")
	Generate(ctx, queries)
	Prompt(ctx, queries, currentUser, currentBatch)
}

func whoAmI(ctx context.Context, currentUser int32, queries *db.Queries, currentBatch int32) {
	user, err := queries.GetUser(ctx, currentUser)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(user)
	Prompt(ctx, queries, currentUser, currentBatch)
}

func listBatches(ctx context.Context, queries *db.Queries, currentUser int32, currentBatch int32) {
	batches, err := queries.ListBatches(ctx)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(batches)
	Prompt(ctx, queries, currentUser, currentBatch)
}

func dispatchedBatches(ctx context.Context, queries *db.Queries, currentUser int32, currentBatch int32) {
	dispatchedBatches, err := queries.ListDispatchedBatches(ctx, true)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(dispatchedBatches)
	Prompt(ctx, queries, currentUser, currentBatch)
}

func listActions(ctx context.Context, queries *db.Queries, currentUser int32, currentBatch int32) {
	actions, err := queries.ListActions(ctx)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(actions)
	Prompt(ctx, queries, currentUser, currentBatch)

}

func listTransactions(ctx context.Context, queries *db.Queries, currentUser int32, currentBatch int32) {
	transactions, err := queries.ListTransactions(ctx)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(transactions)
	Prompt(ctx, queries, currentUser, currentBatch)
}

func startBatch(ctx context.Context, queries *db.Queries, currentUser int32, currentBatch int32) {
	res := ""
	prompt := &survey.Select{
		Message: "<---------What kind of behaviour you want?--------->",
		Options: []string{"Batch all amounts",
			"Choose individual transactions",
		},
	}
	survey.AskOne(prompt, &res)

	switch res {
	case "Batch all amounts":
		batchAllAmounts(ctx, queries, currentUser, currentBatch)
	case "Choose individual transactions":
		fmt.Printf("Not yet implemented, please check again later\n")
		startBatch(ctx, queries, currentUser, currentBatch)

	default:
		Prompt(ctx, queries, currentUser, currentBatch)
	}

}
func batchAllAmounts(ctx context.Context, queries *db.Queries, currentUser int32, currentBatch int32) {
	roundedAmountsChannel := make(chan float64)
	go batchTransactions(ctx, queries, currentUser, roundedAmountsChannel)
	for msg := range roundedAmountsChannel {
		if msg < 100 {
			fmt.Println(msg, " added to the Batch, but not dispatched")
		}
		if msg >= 100 {
			fmt.Printf("Batch sent, new undispatched batch created")
		}
	}
	Prompt(ctx, queries, currentUser, currentBatch)
}

func batchTransactions(ctx context.Context, queries *db.Queries, currentUser int32, roundedAmountsChannel chan float64) {
	actions, _ := queries.ListActions(ctx)
	for _, action := range actions {
		amount := action.Amount
		wholeAmount := int64(amount)
		decimalAmount := amount - float64(wholeAmount)
		roundedAmount := 2 - math.Round(decimalAmount*100)/100
		insertBatchChan := make(chan float64)
		go addToBatch(ctx, queries, currentUser, roundedAmount, insertBatchChan)
		for msg := range insertBatchChan {
			roundedAmountsChannel <- msg
		}

	}
	close(roundedAmountsChannel)
}

func addToBatch(ctx context.Context, queries *db.Queries, currentUser int32, roundedAmount float64, insertBatchChan chan float64) {
	batches, _ := queries.ListBatches(ctx)
	lastCreatedBatch := batches[len(batches)-1]
	batchAmount := lastCreatedBatch.Amount
	amount := batchAmount + roundedAmount
	if batchAmount < 100 {
		queries.UpdateBatch(ctx, db.UpdateBatchParams{ID: lastCreatedBatch.ID, Amount: amount, Dispatched: false})
	} else if batchAmount >= 100 {
		queries.UpdateBatch(ctx, db.UpdateBatchParams{ID: lastCreatedBatch.ID, Amount: amount, Dispatched: true})
		queries.CreateTransaction(ctx, db.CreateTransactionParams{Amount: amount, UserID: currentUser})
		queries.CreateBatch(ctx, db.CreateBatchParams{UserID: currentUser})
	}
	insertBatchChan <- batchAmount
	close(insertBatchChan)
}
