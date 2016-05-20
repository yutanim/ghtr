package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type Repo struct {
	Name        string
	Language    string
	Owner       string
	Description string
	StarNum     string
}

type Result struct {
	repos []Repo
}

func (r *Repo) Setter(name, language, owner, description, starnum string) {
	r.Name = name
	r.Language = language
	r.Owner = owner
	r.Description = description
	r.StarNum = starnum
}

func main() {
	url := createURL("swift", "today")
	results := fetchPage(url)
	fmt.Println(results)
}

func fetchPage(url string) Result {
	res := Result{}
	var repos []Repo
	doc, _ := goquery.NewDocument(url)
	doc.Find(".repo-list-item").Each(func(i int, item *goquery.Selection) {
		var r Repo
		path := strings.TrimSpace(item.Find(".repo-list-name a").Text())
		s := strings.Split(path, "/")
		owner, name := s[0], s[1]
		language := strings.TrimSpace(strings.Split(item.Find("p.repo-list-meta").Text(), "\n")[1])
		starnum := strings.TrimSpace(strings.Split(item.Find("p.repo-list-meta").Text(), "\n")[5])
		desc := strings.TrimSpace(item.Find("p.repo-list-description").Text())
		r.Setter(name, language, owner, desc, starnum)
		repos = append(repos, r)
	})
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
