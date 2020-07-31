package baidu_api

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

func jsonRequest(params []byte, url string) ([]byte, error){
	return request(params, url, "json")
}

func formRequest(params []byte, url string) ([]byte, error) {
	return request(params, url, "form")
}

func request(params []byte, url string, t string) ([]byte, error) {
	var contentType string
	switch strings.ToLower(t) {
	case "json":
		contentType = "application/json;charset=utf-8"
	case "form":
		contentType = "application/x-www-urlencoded"
	default:
		contentType = "application/x-www-urlencoded"
	}

	resp, err := http.Post(url, contentType, bytes.NewBuffer(params))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	return b, nil
}
