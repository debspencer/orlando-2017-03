package goget

import (
	"log"

	"github.com/valyala/fasthttp"
)

func GoGet(url string) []byte {
	var body []byte
	statusCode, body, err := fasthttp.Get(body, url)
	if err != nil {
		log.Fatal("Get Failed", err)
	}
	if statusCode != 200 {
		log.Fatal("Failed to get: ", url, " status:", statusCode)
	}
	return body
}
