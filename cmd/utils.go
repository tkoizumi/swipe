package main

import (
	"strings"
)

func GetUrl(args []string) string {
	for i := 1; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "-") {
			i += 1
		} else {
			return arg
		}
	}
	return ""
}
