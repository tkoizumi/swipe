package main

import (
	"strings"
)

func GetUrl(args []string, flagHasValue map[string]bool) string {
	for i := 1; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "-") {
			flag := []rune(arg)[1]
			if flagHasValue[string(flag)] {
				i += 1
			}
		} else {
			return arg
		}
	}
	return ""
}
