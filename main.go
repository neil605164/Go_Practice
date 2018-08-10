package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	httpGet()
	fmt.Println("Hello World")
}

func httpGet() string {
	resp, err := http.Get("https://tw.yahooooo.com/")
	if err != nil {
		return "HTTP Connect URL ERROR"
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "GET URL INFO ERROR"
	}

	fmt.Println(string(body))
	return ""
}
