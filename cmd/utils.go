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

func MakeFlagMap(flagArr [][]interface{}) map[string]bool {
	flagMap := make(map[string]bool)

	for _, f := range flagArr {
		flagMap[f[0].(string)] = f[1].(bool)
	}
	return flagMap
}
