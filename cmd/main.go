package main

import (
	"fmt"
	"os"
	"strings"
	request "swipe/internal/core/request"
	response "swipe/internal/core/response"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Needs at least url as argument")
		os.Exit(1)
	}

	baseUrl := os.Args[1]
	urlArr := []string{baseUrl}

	queryParams := []string{}

	for i, arg := range os.Args {
		if arg == "-q" && len(os.Args) > i {
			queryParams = append(queryParams, os.Args[i+1])
		}
	}

	if len(queryParams) != 0 {
		urlArr = append(urlArr, "?")
		for _, p := range queryParams {
			urlArr = append(urlArr, p)
			urlArr = append(urlArr, "&")
		}
		urlArr = urlArr[:len(urlArr)-1]
	}

	url := strings.Join(urlArr, "")

	req := request.Create(url)
	r, e := req.Get()

	res := response.Create(r, e)
	res.Print()

}
