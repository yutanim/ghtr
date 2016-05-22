package main

import (
	"fmt"
	"github.com/yutanim/gothubtrends/loader"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	week     = kingpin.Flag("weekly", "Search Weekly").Short('w').Bool()
	month    = kingpin.Flag("monthly", "Search Monthly").Short('m').Bool()
	name     = kingpin.Flag("name", "Input name").Short('n').String()
	language = kingpin.Arg("language", "Input Language").String()
)

func main() {
	kingpin.Parse()
	if *week && *month {
		fmt.Print("Error!!")
		return
	}
	fmt.Printf("verbose mode: %v, count: %d, name: %s\n, name2: %s", *week, *month, *name, *language)
}
