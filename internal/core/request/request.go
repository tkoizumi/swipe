package request

import (
	"fmt"
	"net/http"
	"strings"
	flags "swipe/internal/core/flags"
)

type request struct {
	URL         string
	Base        string
	QueryParams []string
}

func Create(url string, flagArr []flags.Flag) *request {
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
	}
	finalUrl := strings.Join(urlArr, "")

	return &request{URL: finalUrl, Base: url, QueryParams: queryParams}
}

func (r request) Get() (*http.Response, error) {
	fmt.Println("Sending GET request to " + r.URL + "\n")
	res, err := http.Get(r.URL)
	return res, err
}
