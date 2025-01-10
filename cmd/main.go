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
	qFlag := flags.Create("q")
	qFlag.Parse(os.Args)
	xFlag := flags.Create("X")
	xFlag.Parse(os.Args)

	reqFlagArr = append(reqFlagArr, *qFlag)
	reqFlagArr = append(reqFlagArr, *xFlag)

	req := request.Create(baseUrl, reqFlagArr)
	r, e := req.Execute()

	res := response.Create(r, e)
	res.Print()

}
