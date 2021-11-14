package util

import (
	"github.com/AlecAivazis/survey/v2"
)

func Prompt() string {
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
	return res
    }
    