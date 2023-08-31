package api

import (
	"bytes"
	"fmt"
	"github.com/BtQuentin/gitlab-mr-following/pkg/configuration"
	"io"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

type Output struct {
	Summary     string
	Description string
	Image       string
	MRS         string
}

type MRT struct {
	Request     MR
	ProjectName string
}

type MR struct {
	Title     string
	Author    string
	Link      string
	CreatedAt string
	UpdatedAt string
}

func SendMessageToTeams(
	e configuration.Environment,
	teams configuration.Teams,
	mrs []Details,
	pn string,
) {
	mrString := getMessage(mrs)

	sendMessageToTeams(e, teams, mrString, pn)
}

func getMessage(mrs []Details) string {
	var message string
	for i := 0; i < len(mrs); i++ {
		t := template.Must(template.ParseFiles("template/mr.json.tmpl"))

		title := strings.Replace(mrs[i].Title, "\"", "'", -1)
		title = strings.Replace(title, "\\", "/", -1)

		var output bytes.Buffer
		err := t.Execute(&output, MRT{
			Request: MR{
				Title:     title,
				Author:    mrs[i].Author,
				Link:      mrs[i].Link,
				CreatedAt: mrs[i].CreatedAt,
				UpdatedAt: mrs[i].UpdatedAt,
			},
			ProjectName: mrs[i].ProjectName,
		})

		if err != nil {
			panic(err)
		}

		message = message + output.String()

		if i < len(mrs)-1 {
			message = message + ","
		}
	}

	return message
}

func sendMessageToTeams(e configuration.Environment, teams configuration.Teams, mrString, pn string) {
	t := template.Must(template.ParseFiles("template/teams.json.tmpl"))

	desc := e.Description
	if pn != "" {
		desc += " for " + pn + " project "
	}

	var output bytes.Buffer
	err := t.Execute(&output, Output{
		Summary:     e.Summary,
		Description: desc,
		Image:       e.Image,
		MRS:         mrString,
	})

	if err != nil {
		panic(err)
	}

	resp, err := http.Post(teams.API, "application/json", bytes.NewBufferString(output.String()))
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	fmt.Println("\t Return code: " + strconv.Itoa(resp.StatusCode))
	if resp.StatusCode != 200 {
		fmt.Println("\t Response: FAIL")
	} else {
		fmt.Println("\t Response: OK")
	}
}
