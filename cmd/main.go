package main

import (
	"fmt"
	"os"
	"strings"
	flags "swipe/internal/core/flags"
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

	qFlag := flags.Create("q")
	qFlag.Parse(os.Args)

	if len(qFlag.Values) != 0 {
		urlArr = append(urlArr, "?")
		for _, p := range qFlag.Values {
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
