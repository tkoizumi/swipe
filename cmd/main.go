package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Needs at least url as argument")
		os.Exit(1)
	}
	url := os.Args[1]
	fmt.Println("Sending request to " + url)

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Response error")
		os.Exit(1)
	}

	fmt.Println(res)
}
