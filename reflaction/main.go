package main

import (
	"fmt"
	"reflect"
)

type ReachError struct {
	message string
	status  int
}

func main() {

	rErr := ReachError{
		message: "not valid number",
		status:  400,
	}

	type_value := reflect.TypeOf(rErr)

	fmt.Println(type_value.Field(1).Name)
	fmt.Println(type_value.Field(1).Type)

}
