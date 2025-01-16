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
	XFlag := flags.Create("X")
	XFlag.Parse(os.Args)
	HFlag := flags.Create("H")
	HFlag.Parse(os.Args)
	dFlag := flags.Create("d")
	dFlag.Parse(os.Args)
	oFlag := flags.Create("o")
	oFlag.Parse(os.Args)

	reqFlagArr = append(reqFlagArr, *qFlag)
	reqFlagArr = append(reqFlagArr, *XFlag)
	reqFlagArr = append(reqFlagArr, *HFlag)
	reqFlagArr = append(reqFlagArr, *dFlag)

	req := request.Create(baseUrl, reqFlagArr)
	r, e := req.Execute()

	res := response.Create(r, e)
	res.Print()

}
