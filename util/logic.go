package util

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/JoaoLobo94/donut_test/db/sqlc"
	"github.com/davecgh/go-spew/spew"
	"log"
	"math"
)

func Prompt(ctx context.Context, queries *db.Queries, current_user int32, current_batch int32) {
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
		startOver(ctx, queries, current_user, current_batch)
	case "Start Batching low volume bank transactions. This will run concurrently":
		startBatch(ctx, queries, current_user, current_batch)
	case "Which user am I?":
		whoAmI(ctx, current_user, queries, current_batch)
	case "View all batches":
		listBatches(ctx, queries, current_user, current_batch)
	case "View all dispatched batches":
		dispatchedBatches(ctx, queries, current_user, current_batch)
	case "View all transactions sent to the bank":
		listTransactions(ctx, queries, current_user, current_batch)
	case "View all seeded bank user bank transactions to round up":
		listActions(ctx, queries, current_user, current_batch)
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
func startOver(ctx context.Context, queries *db.Queries, current_user int32, current_batch int32) {
	fmt.Printf("Erasing....\n")
	cleanUp(ctx, queries)
	fmt.Printf("Erased \n")
	Generate(ctx, queries)
	Prompt(ctx, queries, current_user, current_batch)
}

func whoAmI(ctx context.Context, current_user int32, queries *db.Queries, current_batch int32) {
	user, err := queries.GetUser(ctx, current_user)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(user)
	Prompt(ctx, queries, current_user, current_batch)
}

func listBatches(ctx context.Context, queries *db.Queries, current_user int32, current_batch int32) {
	batches, err := queries.ListBatches(ctx)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(batches)
	Prompt(ctx, queries, current_user, current_batch)
}

func dispatchedBatches(ctx context.Context, queries *db.Queries, current_user int32, current_batch int32) {
	dispatchedBatches, err := queries.ListDispatchedBatches(ctx, true)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(dispatchedBatches)
	Prompt(ctx, queries, current_user, current_batch)
}

func listActions(ctx context.Context, queries *db.Queries, current_user int32, current_batch int32) {
	actions, err := queries.ListActions(ctx)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(actions)
	Prompt(ctx, queries, current_user, current_batch)

}

func listTransactions(ctx context.Context, queries *db.Queries, current_user int32, current_batch int32) {
	transactions, err := queries.ListTransactions(ctx)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(transactions)
	Prompt(ctx, queries, current_user, current_batch)
}

func startBatch(ctx context.Context, queries *db.Queries, current_user int32, current_batch int32) {
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
		batchAllAmounts(ctx, queries, current_user, current_batch)
	case "Choose individual transactions":
		fmt.Printf("Not yet implemented, please check again later\n")
		startBatch(ctx, queries, current_user, current_batch)

	default:
		Prompt(ctx, queries, current_user, current_batch)
	}

}
func batchAllAmounts(ctx context.Context, queries *db.Queries, current_user int32, current_batch int32) {
	roundedAmountsChannel := make(chan float64)
	go batchTransactions(ctx, queries, current_user,roundedAmountsChannel, current_batch)
	// go addToBatch(ctx, queries, current_user, roundedAmountsMap, current_batch)
	for msg := range roundedAmountsChannel{
		fmt.Println(msg, " added to the batch", current_batch)
	}
	Prompt(ctx, queries, current_user, current_batch)
}

func batchTransactions(ctx context.Context, queries *db.Queries, current_user int32 ,roundedAmountsChannel chan float64, current_batch int32) {
	actions, _ := queries.ListActions(ctx)
	for _, action := range actions {
		amount := action.Amount
		wholeAmount := int64(amount)
		decimalAmount := amount - float64(wholeAmount)
		roundedAmount := 1 - math.Round(decimalAmount*100)/100
		insertBatchChan := make(chan float64)
		// roundedAmountsChannel <- roundedAmount
		go addToBatch(ctx, queries, current_user, roundedAmount, current_batch, insertBatchChan)
		for msg := range insertBatchChan{
			roundedAmountsChannel <- msg
		}

	}
	close(roundedAmountsChannel)
}

func addToBatch(ctx context.Context, queries *db.Queries, current_user int32, roundedAmount float64, current_batch int32, insertBatchChan chan float64) {
	batches, _ := queries.ListBatches(ctx)
	lastCreatedBatch := batches[len(batches) -1]
	batchAmount := lastCreatedBatch.Amount
	amount := batchAmount + roundedAmount
	if batchAmount < 100{
		queries.UpdateBatch(ctx, db.UpdateBatchParams{ID: lastCreatedBatch.ID, Amount: amount, Dispatched: false})
	}else if batchAmount >= 100{
		queries.UpdateBatch(ctx, db.UpdateBatchParams{ID: lastCreatedBatch.ID, Amount: amount, Dispatched: true})
		queries.CreateTransaction(ctx, db.CreateTransactionParams{Amount: amount, UserID: current_user})
		queries.CreateBatch(ctx, db.CreateBatchParams{UserID: current_user})
	}
	insertBatchChan <- batchAmount
	close(insertBatchChan)
}