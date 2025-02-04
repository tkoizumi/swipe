package parser

import (
	"encoding/json"
	"fmt"
)

func ParseJSON(data []byte, fields []string) []string {
	fieldValues := []string{}
	jsonData := []map[string]interface{}{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, field := range fields {
		value, exists := jsonData[0][field]
		if !exists {
			fmt.Println(field, "field not found")
			continue
		}

		if strValue, ok := jsonData[0][field].(string); ok {
			fieldValues = append(fieldValues, strValue)
		} else {
			fieldValues = append(fieldValues, fmt.Sprintf("%v", value))
		}
	}
	return fieldValues
}

func ParseXML(data []byte) map[string]string {
	res := map[string]string{}
	return res
}
