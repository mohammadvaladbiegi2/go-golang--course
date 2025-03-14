package main

import (
	"fmt"
	"strings"
)

func Isogram(word string) bool {
	word = strings.ToLower(word)

	for index, value := range word {
		if string(value) == "-" || string(value) == " " {
			continue
		}
		// بررسی می‌کند که آیا کاراکتر value در زیررشته word از ایندکس index+1 به بعد وجود دارد.
		// اگر وجود داشته باشد، به این معنی است که کلمه ایزوگرام نیست و false برمی‌گرداند.
		if strings.Contains(word[index+1:], string(value)) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Isogram("six-year-old"))
}
