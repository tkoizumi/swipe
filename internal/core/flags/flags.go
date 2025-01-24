package flags

import (
	"fmt"
	"os"
)

type Flag struct {
	Name     string
	HasValue bool
	Values   []string
}

func PrepareAll(args []string, flags [][]interface{}, flagArr *[]Flag) {
	for _, flag := range flags {
		flag := Create(flag[0].(string), flag[1].(bool))
		if flag.HasValue {
			flag.Parse(args)
		}
		*flagArr = append(*flagArr, *flag)
	}
}

func Create(name string, hasValue bool) *Flag {
	return &Flag{Name: name, HasValue: hasValue, Values: []string{}}
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
