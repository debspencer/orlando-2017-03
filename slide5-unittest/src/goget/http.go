package goget

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type HttpInterface interface{
	Get(url string) ([]byte, error)
}

type HttpData struct {
}

func GetHttpInterface() HttpInterface {
	return &HttpData{}
}


func (_ *HttpData) Get(url string) ([]byte, error) {
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
