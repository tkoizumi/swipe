package parser

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func ExtractFields(data []byte, fields []string) []byte {
	parsedData := getValues(data, fields)
	jsonBytes, err := json.MarshalIndent(parsedData, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	return jsonBytes
}

func getValues(data []byte, fields []string) []map[string]string {
	jsonData := []map[string]interface{}{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Error:", err)
	}

	parsedData := []map[string]string{}

	for _, obj := range jsonData {
		fieldValues := map[string]string{}
		for _, field := range fields {

			value, exists := obj[field]
			if !exists {
				fmt.Println(field, "field not found")
				continue
			}

			if strValue, ok := obj[field].(string); ok {
				fieldValues[field] = strValue
			} else {
				fieldValues[field] = fmt.Sprintf("%v", value)
			}
		}
		parsedData = append(parsedData, fieldValues)
	}

	return parsedData
}

func ParseJSON(data []byte, input string) []byte {
	outputArr := []string{}

	variables := parseVar(input)
	fields := createFields(variables)
	varToValueArr := getValues(data, fields)

	for _, varToValue := range varToValueArr {
		output := replaceVars(input, varToValue)
		outputArr = append(outputArr, output)
	}

	outputStr := strings.Join(outputArr[:], "\n")

	return []byte(outputStr)
}

func replaceVars(input string, varToValue map[string]string) string {
	varRegex := regexp.MustCompile(`\.[a-zA-Z_][a-zA-Z0-9_]*`)
	result := varRegex.ReplaceAllStringFunc(input, func(match string) string {
		variable := match[1:]
		if val, exists := varToValue[variable]; exists {
			return val
		}
		return match // Keep original if not found
	})

	return result
}

func parseVar(input string) []string {
	varRegex := regexp.MustCompile(`\.[a-zA-Z_][a-zA-Z0-9_]*`)
	matches := varRegex.FindAllString(input, -1)

	return matches
}

func createFields(variables []string) []string {
	fields := make([]string, len(variables))

	for i, v := range variables {
		fields[i] = v[1:]
	}
	return fields
}

func ParseXML(data []byte) map[string]string {
	res := map[string]string{}
	return res
}

func CreateStruct(struc string) {
	variables := parseVar(struc)

	for _, v := range variables {
		fmt.Println(v)
	}
}
