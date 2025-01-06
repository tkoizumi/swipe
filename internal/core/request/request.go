package request

import (
	"fmt"
	"net/http"
)

type request struct {
	URL string
}

func Create(url string) *request {
	return &request{URL: url}
}

func (r request) Get() (*http.Response, error) {
	fmt.Println("Sending Get request to " + r.URL + "\n")
	res, err := http.Get(r.URL)
	return res, err
}
