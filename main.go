package main

import (
	"github.com/yutanim/ghtr/loader"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	week     = kingpin.Flag("weekly", "Search Weekly Trends").Short('w').Bool()
	month    = kingpin.Flag("monthly", "Search Monthly Trends").Short('m').Bool()
	all      = kingpin.Flag("all", "fetch all trends").Short('a').Bool()
	language = kingpin.Arg("language", "Input Lanaguage which you want to search").String()
)

func main() {
	kingpin.Parse()
	loader.Load(*language, *week, *month, *all)
}
