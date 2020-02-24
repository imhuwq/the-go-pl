package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssueURL = "https://api.github.com/search/issues"

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

func SortSearchResult(searchResult *IssueSearchResult) (map[string][]*Issue, error) {
	result := map[string][]*Issue{
		"BeforeOneYear":  {},
		"BeforeOneMonth": {},
		"WithinOneMonth": {},
	}

	now := time.Now()
	yearAgo := now.AddDate(-1, 0, 0)
	monthAgo := now.AddDate(0, -1, 0)

	for _, issue := range searchResult.Items {
		if issue.CreatedAt.Before(yearAgo) {
			result["BeforeOneYear"] = append(result["BeforeOneYear"], issue)
		} else if issue.CreatedAt.Before(monthAgo) {
			result["BeforeOneMonth"] = append(result["BeforeOneMonth"], issue)
		} else {
			result["WithinOneMonth"] = append(result["WithinOneMonth"], issue)
		}
	}

	return result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues\n\n", result.TotalCount)

	sortedResult, err := SortSearchResult(result)

	for peroid, issues := range sortedResult {
		fmt.Printf("%s:\n", peroid)
		for _, issue := range issues {
			fmt.Printf("#%-5d %s %9.9s %.55s\n", issue.Number, issue.CreatedAt, issue.User.Login, issue.Title)
		}
		fmt.Printf("======\n")
	}
}
