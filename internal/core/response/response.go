package response

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	flags "swipe/internal/core/flags"
)

type response struct {
	Res      *http.Response
	Filename string
	Body     []byte
}

func Create(res *http.Response, flagArr []flags.Flag) *response {
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
	}
	return &response{Res: res, Filename: filename, Body: body}
}

func (r response) Execute() {
	if r.Filename != "" {
		r.Download()
	}
}

func (r response) Print() {
	fmt.Println(string(r.Body))
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
