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

	baseUrl := os.Args[1]
	flagArr := []flags.Flag{}

	qFlag := flags.Create("q")
	qFlag.Parse(os.Args)
	flagArr = append(flagArr, *qFlag)

	req := request.Create(baseUrl, flagArr)
	r, e := req.Get()

	res := response.Create(r, e)
	res.Print()

}
