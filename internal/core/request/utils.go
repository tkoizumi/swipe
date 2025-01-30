package request

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func (r request) Print() {
	fmt.Printf("Sending %s request to %s\n", r.Method, r.URL)
}

func getHeader(header string) []string {
	headerKV := make([]string, 0, 2)

	if strings.Contains(header, "Content-Type") || strings.Contains(header, "SOAPAction") {
		headerKV = strings.Split(header, ":")
		headerKV[1] = strings.TrimSpace(headerKV[1])

	} else {
		headerKV = append(headerKV, "Content-Type")
		headerKV = append(headerKV, header)
	}
	return headerKV
}

func (r request) createNew() *http.Request {
	req, err := http.NewRequest(r.Method, r.URL, r.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return req
}

func setHeaders(headers []string, req *http.Request) {
	for _, header := range headers {
		headerKV := getHeader(header)
		req.Header.Set(headerKV[0], headerKV[1])
	}
}
