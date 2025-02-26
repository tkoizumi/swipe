package request

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"
	flags "swipe/internal/core/flags"
)

type request struct {
	Method         string
	Headers        []string
	URL            string
	Base           string
	QueryParams    []string
	Body           *bytes.Buffer
	Redirect       bool
	User           string
	Password       string
	ClientCertPath string
	PrivateKeyPath string
	TLSConfig      *tls.Config
}

func Create(url string, flagArr []flags.Flag) *request {
	method := "GET" //default method
	body := bytes.NewBuffer(nil)
	redirect := false
	urlArr := []string{url}
	headers := []string{"application/x-www-form-urlencoded"}
	queryParams := []string{}
	user := ""
	password := ""
	clientCertPath := ""
	privateKeyPath := ""

	for _, flag := range flagArr {
		if flag.Name == "X" && len(flag.Values) != 0 {
			method = flag.Values[0]
		}
		if flag.Name == "d" && len(flag.Values) != 0 {
			bodyStr := flag.Values[0]
			if bodyStr[0] == '@' {
				fileContent, err := os.ReadFile(bodyStr[1:])
				if err != nil {
					fmt.Printf("Error reading file: %v\n", err)
					os.Exit(1)
				}
				body = bytes.NewBuffer(fileContent)

			} else {
				body = bytes.NewBuffer([]byte(bodyStr))
			}
		}
		if flag.Name == "H" && len(flag.Values) != 0 {
			headers = []string{}
			for _, header := range flag.Values {
				headers = append(headers, header)
			}
		}
		if flag.Name == "L" {
			redirect = true
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

		if flag.Name == "u" && len(flag.Values) != 0 {
			user = flag.Values[0]
		}
		if flag.Name == "p" && len(flag.Values) != 0 {
			password = flag.Values[0]
		}
		if flag.Name == "E" && len(flag.Values) != 0 {
			clientCertPath = flag.Values[0]
		}
		if flag.Name == "Y" && len(flag.Values) != 0 {
			privateKeyPath = flag.Values[0]
		}
	}
	finalUrl := strings.Join(urlArr, "")

	return &request{
		URL:            finalUrl,
		Headers:        headers,
		Body:           body,
		Base:           url,
		QueryParams:    queryParams,
		Method:         method,
		Redirect:       redirect,
		User:           user,
		Password:       password,
		ClientCertPath: clientCertPath,
		PrivateKeyPath: privateKeyPath,
	}
}

func (r request) Execute() *http.Response {
	res := (*http.Response)(nil)
	err := (error)(nil)

	if isValidMethod(r.Method) {
		r.Print()
		res, err = r.Do()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Error: Invalid or unsupported HTTP method.")
		fmt.Println("Please use a valid HTTP method such as GET, POST, PUT, DELETE, PATCH, etc.")
		os.Exit(1)
	}
	return res
}

func (r request) Do() (*http.Response, error) {
	req, err := http.NewRequest(r.Method, r.URL, r.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	setHeaders(r.Headers, req)
	req.SetBasicAuth(r.User, r.Password)
	r.LoadClientCert()

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: r.TLSConfig,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if r.Redirect {
				return nil
			} else {
				return http.ErrUseLastResponse
			}
		},
	}

	res, err := client.Do(req)

	return res, err
}

func (r *request) LoadClientCert() {
	cert, err := tls.LoadX509KeyPair(r.ClientCertPath, r.PrivateKeyPath)
	if err != nil {
		fmt.Println("Error loading client certificate:", err)
		return
	}

	r.TLSConfig.Certificates = []tls.Certificate{cert}
}

func (r request) Print() {
	fmt.Printf("Sending %s request to %s\n", r.Method, r.URL)
}
