package goget

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func HttpGet(url string) ([]byte, error) {
	var body []byte
	statusCode, body, err := fasthttp.Get(body, url)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, fmt.Errorf("Failed to get: %s status: %d", err.Error(), statusCode)
	}
	return body, nil
}
