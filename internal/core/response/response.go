package response

import (
	"fmt"
	"io"
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
		fmt.Println(r.Err)
		os.Exit(1)
	}

	body, err := io.ReadAll(r.Res.Body)
	if err != nil {
		fmt.Println("Body response error")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(body))

	r.Res.Body.Close()
}
