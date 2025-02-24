package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	flags "swipe/internal/core/flags"

	"github.com/itchyny/gojq"
)

type response struct {
	Res           *http.Response
	Filename      string
	Header        map[string][]string
	Body          []byte
	IncludeHeader bool
	ParseStruct   string
}

func Create(res *http.Response, flagArr []flags.Flag) *response {
	includeHeader := false
	header := map[string][]string{}
	parseStruct := ""

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
			parseStruct = flag.Values[0]
		}
	}
	return &response{
		Res:           res,
		Filename:      filename,
		Header:        header,
		Body:          body,
		IncludeHeader: includeHeader,
		ParseStruct:   parseStruct,
	}
}

func (r response) Execute() {
	if r.ParseStruct != "" {
		r.Parse()
	}
	r.Print()
	if r.Filename != "" {
		r.Download()
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

func (r *response) Parse() {
	content_type := r.Header["Content-Type"][0]
	format := detectFormat(content_type)
	if format == "" {
		fmt.Println("Malformed data")
		os.Exit(1)
	}
	if format == "json" {
		var jsonData interface{}

		err := json.Unmarshal(r.Body, &jsonData)
		if err != nil {
			fmt.Println("Failed to parse JSON:", err)
			os.Exit(1)
		}

		query, err := gojq.Parse(r.ParseStruct)
		if err != nil {
			fmt.Println("Invalid JQ query:", err)
			os.Exit(1)
		}

		iter := query.Run(jsonData)
		var results []interface{}
		for {
			value, ok := iter.Next()
			fmt.Println("value: ", value)
			if !ok {
				break
			}
			if err, isErr := value.(error); isErr {
				fmt.Println("JQ Error:", err)
				os.Exit(1)
			}
			results = append(results, value)
		}

		jsonBytes, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			os.Exit(1)
		}

		r.Body = jsonBytes
	}
}
