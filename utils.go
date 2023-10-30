package main

import (
	"fmt"
	"reflect"
)

func CheckEmptyKeyAndValue(valuesObject reflect.Value) []string {
	missingFields := []string{}
	typeObject := valuesObject.Type()

	if valuesObject.Kind() == reflect.Ptr && typeObject.Kind() == reflect.Ptr {
		valuesObject = valuesObject.Elem()
		typeObject = typeObject.Elem()
	}

	if valuesObject.Kind() == reflect.Struct {
		for i := 0; i < valuesObject.NumField(); i++ {
			nameField := typeObject.Field(i).Name
			value := valuesObject.Field(i).String()
			if value == "" {
				message := fmt.Sprintf("Field %v is missing", nameField)
				missingFields = append(missingFields, message)
			}
		}
	}

	if len(missingFields) == 0 {
		return nil
	}

	return missingFields
}
