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

	baseUrl := GetUrl(os.Args)

	reqFlagArr := []flags.Flag{}
	resFlagArr := []flags.Flag{}

	reqFlags := []string{"q", "X", "H", "d"}
	flags.PrepareAll(os.Args, reqFlags, &reqFlagArr)

	resFlags := []string{"o"}
	flags.PrepareAll(os.Args, resFlags, &resFlagArr)

	req := request.Create(baseUrl, reqFlagArr)
	r := req.Execute()

	res := response.Create(r, resFlagArr)
	res.Print()
	res.Execute()
}
