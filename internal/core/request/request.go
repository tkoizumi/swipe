package request

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
	flags "swipe/internal/core/flags"
)

type request struct {
	Method      string
	Header      string
	URL         string
	Base        string
	QueryParams []string
	Body        *bytes.Buffer
}

func Create(url string, flagArr []flags.Flag) *request {
	method := "GET" //default method
	header := "application/x-www-form-urlencoded"
	body := bytes.NewBuffer(nil)

	urlArr := []string{url}
	queryParams := []string{}

	for _, flag := range flagArr {
		if flag.Name == "X" && len(flag.Values) != 0 {
			method = flag.Values[0]
		}
		if flag.Name == "H" && len(flag.Values) != 0 {
			header = flag.Values[0]
		}
		if flag.Name == "d" && len(flag.Values) != 0 {
			bodyStr := flag.Values[0]
			body = bytes.NewBuffer([]byte(bodyStr))
		}

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

	return &request{
		URL:         finalUrl,
		Header:      header,
		Body:        body,
		Base:        url,
		QueryParams: queryParams,
		Method:      method,
	}
}

func (r request) Execute() (*http.Response, error) {
	res := (*http.Response)(nil)
	err := (error)(nil)

	if r.Method == "GET" {
		res, err = r.Get()
	} else if r.Method == "POST" {
		res, err = r.Post()
	} else if r.Method == "PUT" {
		res, err = r.Put()

	} else {
		fmt.Println("Error: Invalid or unsupported HTTP method.")
		fmt.Println("Please use a valid HTTP method such as GET, POST, PUT, DELETE, PATCH, etc.")
		os.Exit(1)
	}
	return res, err
}

func (r request) Get() (*http.Response, error) {
	fmt.Println("Sending GET request to ", r.URL)
	res, err := http.Get(r.URL)
	return res, err
}

func (r request) Post() (*http.Response, error) {
	header := r.Header
	body := r.Body

	fmt.Println("Sending POST request to ", r.URL)
	fmt.Println("header: ", header)
	fmt.Println("body: ", body)

	res, err := http.Post(r.URL, r.Header, r.Body)
	return res, err
}

func (r request) Put() (*http.Response, error) {
	header := strings.Split(r.Header, " ")
	headerKey := header[0][:len(header[0])-1]
	body := r.Body

	fmt.Println("Sending PUT request to ", r.URL)
	fmt.Println("header: ", r.Header)
	fmt.Println("body: ", body)

	req, err := http.NewRequest(http.MethodPut, r.URL, r.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	req.Header.Set(headerKey, header[1])

	client := &http.Client{}
	res, err := client.Do(req)

	return res, err
}
