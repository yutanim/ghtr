package main

import (
	//"fmt"
	"github.com/PuerkitoBio/goquery"
	//"io"
)

type Repo struct {
	Name        string
	Language    string
	Owner       string
	Description string
	StarNum     int
}

type Result struct {
	repos []Repo
}

func (r *Repo) Setter(name, language, owner, description string, starnum int) {
	r.Name = name
	r.Language = language
	r.Owner = owner
	r.Description = description
	r.StarNum = starnum
}

func main() {
	url := createURL("swift", "today")
	results := fetchPage(url)
}

func fetchPage(url string) Result {
	res := Result{}
	var repos []Repo
	doc, _ := goquery.NewDocument(url)

	res.repos = repos
	return res
}

func createURL(language, since string) string {
	url := "https://github.com/trending"
	if language != "" {
		url += "/" + language
	}
	if since != "" {
		url += "?since=" + since
	}

	return url
}
