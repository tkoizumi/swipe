package response

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	flags "swipe/internal/core/flags"
	parser "swipe/internal/core/parser"
)

type response struct {
	Res           *http.Response
	Filename      string
	Header        map[string][]string
	Body          []byte
	IncludeHeader bool
	ParseFields   []string
}

func Create(res *http.Response, flagArr []flags.Flag) *response {
	includeHeader := false
	header := map[string][]string{}
	parseFields := []string{}

	for k, values := range res.Header {
		v := []string{}
		for _, value := range values {
			v = append(v, value)
		}
		header[k] = values
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Err:", err)
		os.Exit(1)
	}
	res.Body.Close()

	filename := ""
	for _, flag := range flagArr {
		if flag.Name == "o" && len(flag.Values) != 0 {
			filename = flag.Values[0]
		}
		if flag.Name == "i" && flag.InArg {
			includeHeader = true
		}
		if flag.Name == "P" && len(flag.Values) != 0 {
			parseFields = flag.Values
		}
	}
	return &response{Res: res, Filename: filename, Header: header, Body: body, IncludeHeader: includeHeader, ParseFields: parseFields}
}

func (r response) Execute() {
	r.Print()
	if r.Filename != "" {
		r.Download()
	}
	if len(r.ParseFields) != 0 {
		r.Parse()
	}
}

func (r response) Print() {
	if r.IncludeHeader {
		for k, v := range r.Header {
			fmt.Print(k + ": ")
			fmt.Println(v)
		}
	}
	fmt.Println(string(r.Body))
	fmt.Println()
}

func (r response) Download() {
	bodyReader := bytes.NewReader(r.Body)

	file, err := os.OpenFile(r.Filename, os.O_CREATE|os.O_WRONLY, 0644) // Open for reading
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	_, err = io.Copy(file, bodyReader)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	defer file.Close()

	fmt.Println("Response saved as", r.Filename)
}

func (r response) Parse() {
	content_type := r.Header["Content-Type"][0]
	format := detectFormat(content_type)
	if format == "" {
		fmt.Println("Malformed data")
		os.Exit(1)
	}
	if format == "json" {
		fmt.Println(parser.ParseJSON(r.Body, r.ParseFields))
	}
}

func detectFormat(content_type string) string {
	if strings.Contains(content_type, "application/json") {
		return "json"
	}
	if strings.Contains(content_type, "text/xml") {
		return "xml"
	}
	return ""
}
