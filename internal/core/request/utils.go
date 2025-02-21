package request

import (
	"net/http"
	"strings"
)

func getHeader(header string) []string {
	headerKV := make([]string, 0, 2)

	if strings.Contains(header, ":") {
		headerKV = strings.Split(header, ":")
		headerKV[1] = strings.TrimSpace(headerKV[1])

	} else {
		headerKV = append(headerKV, "Content-Type")
		headerKV = append(headerKV, header)
	}
	return headerKV
}

func setHeaders(headers []string, req *http.Request) {
	for _, header := range headers {
		headerKV := getHeader(header)
		req.Header.Set(headerKV[0], headerKV[1])
	}
}
