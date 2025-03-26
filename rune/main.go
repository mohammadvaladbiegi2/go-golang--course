package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "محمد"

	fmt.Println(utf8.RuneCountInString(name))
}
