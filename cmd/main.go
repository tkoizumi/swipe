package main

import (
	"fmt"
	"os"
	flags "swipe/internal/core/flags"
	request "swipe/internal/core/request"
	response "swipe/internal/core/response"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Needs at least url as argument")
		os.Exit(1)
	}

	reqFlagArr := []flags.Flag{}
	resFlagArr := []flags.Flag{}

	// flag id, does flag have value
	reqFlags := [][]interface{}{
		{"q", true},
		{"X", true},
		{"H", true},
		{"d", true},
		{"L", false},
		{"u", true},
	}

	flags.PrepareAll(os.Args, reqFlags, &reqFlagArr)

	resFlags := [][]interface{}{
		{"o", true},
		{"i", false},
	}
	flags.PrepareAll(os.Args, resFlags, &resFlagArr)

	flagMap := make(map[string]bool)

	for _, reqFlag := range reqFlags {
		flagMap[reqFlag[0].(string)] = reqFlag[1].(bool)
	}
	for _, resFlag := range reqFlags {
		flagMap[resFlag[0].(string)] = resFlag[1].(bool)
	}

	baseUrl := GetUrl(os.Args, flagMap)

	req := request.Create(baseUrl, reqFlagArr)
	r := req.Execute()

	res := response.Create(r, resFlagArr)
	res.Print()
	res.Execute()
}
