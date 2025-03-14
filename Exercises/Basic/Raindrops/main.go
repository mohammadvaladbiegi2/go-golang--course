package main

import (
	"fmt"
	"strings"
)

func Raindrops(count int) string {
	var result strings.Builder

	if count%3 == 0 {
		result.WriteString("Pling")
	}
	if count%5 == 0 {
		result.WriteString("Plang")
	}
	if count%7 == 0 {
		result.WriteString("Plong")
	}

	if result.Len() == 0 {
		return fmt.Sprintf("%d", count) // اگر هیچ کدام از شرایط برقرار نبود، عدد را برگردانید
	}

	return result.String() // نتیجه نهایی را برگردانید
}

func main() {
	fmt.Println(Raindrops(30)) // مثال: 28 را به تابع Raindrops ارسال کنید
}
