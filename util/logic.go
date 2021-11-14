package util

import (
	"fmt"
	"context"
	"github.com/JoaoLobo94/donut_test/db/sqlc"
	"github.com/AlecAivazis/survey/v2"
)

func Prompt(ctx context.Context, queries *db.Queries) {
	res := ""
	prompt := &survey.Select{
	    Message: "What would you like to do now?",
	    Options: []string{"Start Over --> This will erase all database",
	    "Start Batching all low volume bank transactions. This will run concurrently",
	    "View all active users",
	    "View all batches",
	    "View all dispatched batches",
	    "View all undispatched batches",
	    "View all transactions sent to the bank"},
	}
	survey.AskOne(prompt, &res) 
	fmt.Printf("Ok, lets: " + res)

	switch res {
	case "Start Over --> This will erase all database":
		startOver(ctx, queries)
	case "Start Batching all low volume bank transactions. This will run concurrently":
		startBatch(ctx, queries)
	case "View all active users":
		// listUsers()
	case "View all batches":
		// listBatches()
	case "View all dispatched batches":
		// listDispatchedBatches()
	case "View all transactions sent to the bank":
		// listTransactions()
	default:
		fmt.Println("Exiting")
	}
    }
    
    func startOver(ctx context.Context, queries *db.Queries) {
	fmt.Printf("Erasing....")
	queries.DeleteActions(ctx)
	queries.DeleteBatches(ctx)
	queries.DeleteTransactions(ctx)
	queries.DeleteUsers(ctx)
	fmt.Printf("Erased")
	Prompt(ctx, queries)
    }

    func startBatch(ctx context.Context, queries *db.Queries) {
	    
    }