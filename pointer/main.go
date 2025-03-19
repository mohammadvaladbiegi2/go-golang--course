package main

import "fmt"

func main() {
	//  & ادرس محل ذخیره
	//  * مقدار که در اون ادرس ذخیره شده

	a := 55
	b := &a
	fmt.Printf("addrese of a ==> %p\n addres of b ==> %p ", &a, b)
	fmt.Println("value of B", *b)
	*b = 60 // ادرس "بی" همون "ا" هست ولی مقدارش اینجا تغییر میدیم

}
