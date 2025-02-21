package response

import (
	"strings"
)

func detectFormat(content_type string) string {
	if strings.Contains(content_type, "application/json") {
		return "json"
	}
	if strings.Contains(content_type, "text/xml") {
		return "xml"
	}
	return ""
}
