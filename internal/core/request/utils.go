package request

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func (r request) Print() {
	fmt.Printf("Sending %s request to %s\n", r.Method, r.URL)
	fmt.Println("header: ", r.Header)
	fmt.Println("body: ", r.Body)
}

func (r request) getHeader() string {
	header := r.Header
	if strings.Contains(header, "Content-Type") {
		header = strings.Split(r.Header, " ")[1]
	}
	return header
}

func (r request) createNew() *http.Request {
	req, err := http.NewRequest(r.Method, r.URL, r.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return req
}
