package response

import (
	"fmt"
	"net/http"
	"os"
)

type response struct {
	Res *http.Response
	Err error
}

func Create(res *http.Response, err error) *response {
	return &response{Res: res, Err: err}
}

func (r response) Print() {
	if r.Err != nil {
		fmt.Println("Response error")
		os.Exit(1)
	}
	fmt.Println(r.Res)
}
