package configuration

import (
	"flag"
	"fmt"
	"os"
)

type Flags struct {
	Token   string
	Webhook string
	Team    string
	Project int
}

func LoadFlags() Flags {
	var (
		tokenPtr   = flag.String("token", "", "Gitlab token (required)")
		webhookPtr = flag.String("webhook", "", "Webhook URL for Teams (required)")
		teamPtr    = flag.String("team", "", "Your team name (Example: SOP)")
		projectPtr = flag.Int("project", -1, "Gitlab project ID (Example: 9 for SophiaHTML)")
	)
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("No arguments")
		fmt.Println("Try lunching with -h to show how to use")
		os.Exit(3)
	}

	if *tokenPtr == "" || *webhookPtr == "" {
		fmt.Println("Error")
		fmt.Println("Token and Webhook are required")
		os.Exit(3)
	}

	if *teamPtr != "" && *projectPtr != -1 {
		fmt.Println("Error")
		fmt.Println("You can't use both 'team' and 'project' arguments")
		os.Exit(3)
	}

	if *tokenPtr != "" && *webhookPtr != "" && (*teamPtr == "" && *projectPtr == -1) {
		fmt.Println("Error")
		fmt.Println("You need to set at least on of 'team' or 'project' argument")
		os.Exit(3)
	}

	return Flags{
		Token:   *tokenPtr,
		Webhook: *webhookPtr,
		Team:    *teamPtr,
		Project: *projectPtr,
	}
}
