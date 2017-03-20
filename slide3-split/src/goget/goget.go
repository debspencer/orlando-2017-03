package goget

import (
	"log"
)

func GoGet(url string) []byte {
	body, err := HttpGet(url)
	if err != nil {
		log.Fatal("Get Failed:", err)
	}
	return body
}
