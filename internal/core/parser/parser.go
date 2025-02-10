package parser

import (
	"encoding/json"
	"fmt"
)

func ParseJSON(data []byte, fields []string) []byte {
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

	jsonBytes, err := json.MarshalIndent(parsedData, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	return jsonBytes
}

func ParseXML(data []byte) map[string]string {
	res := map[string]string{}
	return res
}
