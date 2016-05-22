package main

import (
	"github.com/yutanim/gothubtrends/loader"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	week     = kingpin.Flag("weekly", "Search Weekly").Short('w').Bool()
	month    = kingpin.Flag("monthly", "Search Monthly").Short('m').Bool()
	language = kingpin.Arg("language", "Input Language").String()
)

func main() {
	kingpin.Parse()
	loader.Load(*language, *week, *month)
}
