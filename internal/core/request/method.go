package request

var httpMethods = map[string]bool{
	"GET":     true,
	"POST":    true,
	"PUT":     true,
	"DELETE":  true,
	"PATCH":   true,
	"HEAD":    true,
	"OPTIONS": true,
	"CONNECT": true,
	"TRACE":   true,
}

func isValidMethod(method string) bool {
	return httpMethods[method]
}
