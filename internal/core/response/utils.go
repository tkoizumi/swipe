package response

import (
	"fmt"
	"os"
	"strings"
	parser "swipe/internal/core/parser"
)

func (r *response) Extract() {
	content_type := r.Header["Content-Type"][0]
	format := detectFormat(content_type)
	if format == "" {
		fmt.Println("Malformed data")
		os.Exit(1)
	}
	if format == "json" {
		jsonBytes := parser.ExtractFields(r.Body, r.ParseFields)
		r.Body = jsonBytes
	}
}

func (r *response) Parse() {
	content_type := r.Header["Content-Type"][0]
	format := detectFormat(content_type)
	if format == "" {
		fmt.Println("Malformed data")
		os.Exit(1)
	}
	if format == "json" {
		jsonBytes := parser.ParseJSON(r.Body, r.ParseStruct)
		r.Body = jsonBytes
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
