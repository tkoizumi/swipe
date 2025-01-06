package request

import (
	"net/http"
)

type request struct {
	URL string
}

func Create(url string) *request {
	return &request{URL: url}
}

func (r request) Get() (*http.Response, error) {
	res, err := http.Get(r.URL)
	return res, err
}
