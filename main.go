package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "http://github-trends.ryotarai.info/rss/github_trends"

	hoge := "hogee"
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	fmt.Println(hoge)
}
