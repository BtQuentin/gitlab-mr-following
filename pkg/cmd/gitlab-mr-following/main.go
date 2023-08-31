package main

import (
	"fmt"
	"github.com/BtQuentin/gitlab-mr-following/pkg/api"
	"github.com/BtQuentin/gitlab-mr-following/pkg/configuration"
)

func main() {
	flags := configuration.LoadFlags()
	env := configuration.LoadEnv()

	fmt.Println("Connection to Gitlab ...")
	gitlab := configuration.InitGitlab(flags, env)

	fmt.Println("OK. Now, getting MR")

	var mrs []api.Details
	var pn string

	if flags.Project != -1 {
		fmt.Println("Search for whole projects")
		mrs = api.GetMergeRequestsByProject(gitlab, flags.Project, "created_at", "desc", "no")
		pn = api.GetProjectName(gitlab, flags.Project)

	} else {
		fmt.Println("Search for team " + flags.Team)
		mrs = api.GetMergeRequestsByTeam(gitlab, flags.Team, "created_at", "asc", "no")
		pn = ""
	}

	fmt.Println("OK. Now, connection to teams ...")
	teams := configuration.InitTeam(flags)
	fmt.Println("OK. Sending message to teams ...")
	api.SendMessageToTeams(env, teams, mrs, pn)
}
