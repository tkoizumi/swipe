package request

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	flags "swipe/internal/core/flags"
)

type request struct {
	URL         string
	Base        string
	QueryParams []string
	Method      string
}

func Create(url string, flagArr []flags.Flag) *request {
	method := "GET" //default method
	urlArr := []string{url}
	queryParams := []string{}

	for _, flag := range flagArr {
		if flag.Name == "q" {
			urlArr = append(urlArr, "?")

			for _, queryParam := range flag.Values {
				urlArr = append(urlArr, queryParam)
				urlArr = append(urlArr, "&")
				queryParams = append(queryParams, queryParam)
			}
			urlArr = urlArr[:len(urlArr)-1]
		}
		if flag.Name == "X" && len(flag.Values) != 0 {
			method = flag.Values[0]
		}
	}
	finalUrl := strings.Join(urlArr, "")

	return &request{URL: finalUrl, Base: url, QueryParams: queryParams, Method: method}
}

func (r request) Execute() (*http.Response, error) {
	res := (*http.Response)(nil)
	err := (error)(nil)

	if r.Method == "GET" {
		res, err = r.Get()
	} else {
		fmt.Println("Error: Invalid or unsupported HTTP method.")
		fmt.Println("Please use a valid HTTP method such as GET, POST, PUT, DELETE, PATCH, etc.")
		os.Exit(1)
	}
	return res, err
}

func (r request) Get() (*http.Response, error) {
	fmt.Println("Sending GET request to " + r.URL + "\n")
	res, err := http.Get(r.URL)
	return res, err
}
