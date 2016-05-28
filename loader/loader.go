package loader

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"strings"
)

func show(rs []Repositry) {
	for i, v := range rs {
		c := color.New(color.FgCyan, color.Bold)
		c.Printf("#%v  %v/%v (%v)   ", i+1, v.Owner, v.Name, v.Language)
		fmt.Println(v.StarNum)
		fmt.Printf("Description:%v \nURL: https://github.com%v   \n", v.Description, v.URL)
		fmt.Println("--------------------------------------------------------------------")
	}
}

func Load(lang string, weekly, monthly bool) {
	url, err := createURL(lang, weekly, monthly)
	if err != nil {
		fmt.Print(err)
		return
	}
	results := fetchPage(url)
	show(results[:10])
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

func createURL(language string, weekly, monthly bool) (string, error) {
	if weekly && monthly {
		return "", errors.New("Error:You can't set both weekly and monthly")
	}
	url := "https://github.com/trending"
	if language != "" {
		url += "/" + language
	}
	if weekly {
		url += "?since=weekly"
	} else if monthly {
		url += "?since=monthly"
	} else {
		url += "?since=daily"
	}
	return url, nil
}
