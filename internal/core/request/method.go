package request

var httpMethods = map[string]bool{
	"GET":    true,
	"POST":   true,
	"PUT":    true,
	"DELETE": true,
}

func isValidMethod(method string) bool {
	return httpMethods[method]
}
