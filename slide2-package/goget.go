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
	fmt.Println(goget.GoGet(args[0]))
}
