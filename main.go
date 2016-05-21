package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type Repositry struct {
	Name        string
	Language    string
	Owner       string
	Description string
	URL         string
	StarNum     string
}

func (r *Repositry) Setter(name, language, owner, description, url, starnum string) {
	r.Name = name
	r.Language = language
	r.Owner = owner
	r.Description = description
	r.URL = url
	r.StarNum = starnum
}

func main() {
	url := createURL("swift", "today")
	results := fetchPage(url)
	fmt.Println(results)
}

func parse(item *goquery.Selection) Repositry {
	var r Repositry
	repoURL, _ := item.Find(".repo-list-name a").Attr("href")
	s := strings.Split(repoURL, "/")
	_, owner, name := s[0], s[1], s[2]
	language := strings.TrimSpace(strings.Split(item.Find("p.repo-list-meta").Text(), "\n")[1])
	starnum := strings.TrimSpace(strings.Split(item.Find("p.repo-list-meta").Text(), "\n")[5])
	desc := strings.TrimSpace(item.Find("p.repo-list-description").Text())
	r.Setter(name, language, owner, desc, repoURL, starnum)
	return r
}

func fetchPage(url string) []Repositry {
	var res []Repositry
	doc, _ := goquery.NewDocument(url)
	doc.Find(".repo-list-item").Each(func(i int, item *goquery.Selection) {
		res = append(res, parse(item))
	})
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
