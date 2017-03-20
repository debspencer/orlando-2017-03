package main

import 	"flag"
import 	"fmt"
import 	"log"
import 	"github.com/valyala/fasthttp"

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("usage: goget <url>")
	}
	var body []byte
	statusCode, body, err := fasthttp.Get(body, args[0])
	if err != nil {
		log.Fatal("Get Failed", err)
	}
	if statusCode == 200 {
		fmt.Println(string(body))
	} else {
		log.Fatal("Failed to get: ", args[0], " status:", statusCode)
	}
}
