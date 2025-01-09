package flags

import (
	"fmt"
	"os"
)

type Flag struct {
	Name   string
	Values []string
}

func Create(name string) *Flag {
	return &Flag{Name: name, Values: []string{}}
}

func (f *Flag) Parse(args []string) {
	queryParams := []string{}
	flag := "-" + f.Name

	for i, arg := range args {
		if arg == flag {
			if len(args) > (i + 1) {
				queryParams = append(queryParams, args[i+1])
			} else {
				fmt.Println("Missing flag values")
				os.Exit(1)
			}
		}
	}

	f.Values = queryParams
}
