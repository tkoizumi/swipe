package request

import (
	"fmt"
	"net/http"
	"os"
)

type request struct {
	URL string
}

func Create(url string) *request {
	return &request{URL: url}
}

func (r request) Get() *http.Response {
	res, err := http.Get(r.URL)

	if err != nil {
		fmt.Println("Response error")
		os.Exit(1)
	}
	return res
}
