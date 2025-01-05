package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Needs at least url as argument")
		os.Exit(1)
	}
	url := os.Args[1]
	fmt.Println("Sending request to " + url)
}
