package main

import (
	"github.com/imhuwq/the-go-pl/ch04/tutorial/github"
	"html/template"
	"log"
	"os"
	"time"
)

func DaysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

const temp1 = `
<h1>{{.TotalCount}} issues</h1>
<table>
	<tr style='text-align: left'>
		<th>#</th>
		<th>State</th>
		<th>User</th>
		<th>Title</th>
	</tr>
	
{{range .Items}}
	<tr>
		<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
{{end}}
</table>
`

var report = template.Must(template.New("issuehtml").
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
