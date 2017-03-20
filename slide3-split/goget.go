package main

import (
	"flag"
	"fmt"
	"log"

	"goget"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("usage: goget <url>")
	}
	url := args[0]
	body := goget.GoGet(url)
	if len(body) > 0 {
		fmt.Println(string(body))
	}
}
