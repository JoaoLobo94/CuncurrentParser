package util

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/JoaoLobo94/donut_test/db/sqlc"
	"github.com/davecgh/go-spew/spew"
	"log"
	"math"
	// "sync"
)

func Prompt(ctx context.Context, queries *db.Queries, current_user int32, current_batch int32) {
	res := ""
	prompt := &survey.Select{
		Message: "<---------What would you like to do now?--------->",
		Options: []string{"Start Over --> This will erase all database",
			"Start Batching low volume bank transactions. This will run concurrently",
			"Which user am I?",
			"View all batches",
			"View all dispatched batches",
			"View all transactions sent to the bank",
			"View all low volume transactions in the bank"},
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
	case "View all low volume transactions in the bank":
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
	case "Choose individual transactions?":
	// batchSpecificTransactions(ctx, current_user, queries)
	default:
		Prompt(ctx, queries, current_user, current_batch)
	}

}
func batchAllAmounts(ctx context.Context, queries *db.Queries, current_user int32, current_batch int32) {
	roundedAmountsMap := roundUpAllActions(ctx, queries)
	addToBatch(ctx, queries, current_user, roundedAmountsMap, current_batch)
	Prompt(ctx, queries, current_user, current_batch)
}

func roundUpAllActions(ctx context.Context, queries *db.Queries) map[int32]float64 {
	actions, err := queries.ListActions(ctx)
	if err != nil {
		log.Fatal(err)
	}

	roundedAmountsMap := make(map[int32]float64)
	for _, action := range actions {
		amount := action.Amount
		wholeAmount := int64(amount)
		decimalAmount := amount - float64(wholeAmount)
		roundedUpAmounts := 1 - math.Round(decimalAmount*100)/100
		roundedAmountsMap[action.ID] = roundedUpAmounts

	}
	return roundedAmountsMap
}

func addToBatch(ctx context.Context, queries *db.Queries, current_user int32, roundedAmountsMap map[int32]float64, current_batch int32) {
	valueForBatch := 0.0
	loopCounter := 0
	for _, amount := range roundedAmountsMap {
		loopCounter++
		valueForBatch += amount
		if valueForBatch >= 100 {
			current_batch = transactWhenBatchFull(ctx, queries,current_user, current_batch, valueForBatch)
			valueForBatch = 0.0
		} else if loopCounter == len(roundedAmountsMap) && valueForBatch < 100 {
			queries.UpdateBatch(ctx, db.UpdateBatchParams{ID: current_batch, Amount: valueForBatch, Dispatched: false})
			fmt.Printf("There is at least one undisptached batch in the DB.")
		}
	}
}

func transactWhenBatchFull(ctx context.Context, queries *db.Queries, current_user int32, current_batch int32, valueForBatch float64) int32{
	queries.UpdateBatch(ctx, db.UpdateBatchParams{ID: current_batch ,Amount: valueForBatch, Dispatched: true})
	queries.CreateTransaction(ctx, db.CreateTransactionParams{Amount: valueForBatch, UserID: current_user})
	newBatch, err := queries.CreateBatch(ctx, db.CreateBatchParams{UserID: current_user})
	if err != nil {
		log.Fatal(err)
	}
	return newBatch.ID
}
