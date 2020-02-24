package main

import (
	"github.com/imhuwq/the-go-pl/ch04/tutorial/github"
	"log"
	"os"
	"text/template"
	"time"
)

func DaysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

const temp1 = `{{.TotalCount}} Issues:
{{range .Items}}
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s" }}
Age:    {{.CreatedAt | DaysAgo}} Days
{{end}}
`

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"DaysAgo": DaysAgo}).
	Parse(temp1))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	err = report.Execute(os.Stdout, result)
	if err != nil {
		log.Fatal(err)
	}
}
