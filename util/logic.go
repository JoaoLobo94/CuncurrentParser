package util

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/JoaoLobo94/donut_test/db/sqlc"
	"github.com/davecgh/go-spew/spew"
	"log"
	// "sync"
)

func Prompt(ctx context.Context, queries *db.Queries, current_user int32) {
	res := ""
	prompt := &survey.Select{
		Message: "<---------What would you like to do now?--------->",
		Options: []string{"Start Over --> This will erase all database",
			"Start Batching all low volume bank transactions. This will run concurrently",
			"Which user am I?",
			"View all batches",
			"View all dispatched batches",
			"View all transactions sent to the bank",
			"View all low volume transactions that were seeded"},
	}
	survey.AskOne(prompt, &res)

	switch res {
	case "Start Over --> This will erase all database":
		startOver(ctx, queries, current_user)
	case "Start Batching all low volume bank transactions. This will run concurrently":
		startBatch(ctx, queries, current_user)
	case "Which user am I?":
		whoAmI(ctx, current_user, queries)
	case "View all batches":
		listBatches(ctx, queries, current_user)
	case "View all dispatched batches":
		dispatchedBatches(ctx, queries, current_user)
	case "View all transactions sent to the bank":
		listTransactions(ctx, queries,current_user)
	case "View all low volume transactions":
		listActions(ctx, queries, current_user)
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
func startOver(ctx context.Context, queries *db.Queries, current_user int32) {
	fmt.Printf("Erasing....")
	cleanUp(ctx, queries)
	fmt.Printf("Erased -->")
	Generate(ctx, queries)
	Prompt(ctx, queries, current_user)
}

func whoAmI(ctx context.Context, current_user int32, queries *db.Queries) {
	user, err := queries.GetUser(ctx, current_user)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(user)
	Prompt(ctx, queries, current_user)
}

func listBatches(ctx context.Context, queries *db.Queries, current_user int32) {
	batches, err := queries.ListBatches(ctx)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(batches)
	Prompt(ctx, queries, current_user)
}

func dispatchedBatches(ctx context.Context, queries *db.Queries, current_user int32) {
	dispatchedBatches, err := queries.ListDispatchedBatches(ctx, true)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(dispatchedBatches)
	Prompt(ctx, queries, current_user)
}

func listActions(ctx context.Context, queries *db.Queries, current_user int32) {
	actions, err := queries.ListActions(ctx)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(actions)

}

func listTransactions(ctx context.Context, queries *db.Queries, current_user int32) {
	transactions, err := queries.ListTransactions(ctx)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(transactions)
	Prompt(ctx, queries, current_user)
}

func startBatch(ctx context.Context, queries *db.Queries, current_user int32) {
	fmt.Printf("How much would you like to send to the bank? Keep in mind the more bank actions you generated the more accurate will be the number\n")
	fmt.Printf("Number must be above 100 so we don't spam your bank account\n")
	var ammountRequested string
	fmt.Scanln(&ammountRequested)
	// var wg sync.WaitGroup
	// go roundupAllActions(ctx, queries, current_user)
	// go addToBatch()
}
