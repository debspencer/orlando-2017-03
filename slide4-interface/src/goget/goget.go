package goget

import (
	"log"
)

var (
	httpInterface = GetHttpInterface()
)

func GoGet(url string) []byte {
	body, err := httpInterface.Get(url)
	if err != nil {
		log.Fatal("Get Failed:", err)
	}
	return body
}
