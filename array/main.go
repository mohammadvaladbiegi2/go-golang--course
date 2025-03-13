package main

import "fmt"

func main() {

	var array1 [6]int

	for index, value := range array1 {
		fmt.Println(index, value)
	}

	//  Slice

	var slice1 []int

	slice1 = append(slice1, 1, 2, 3, 4, 5, 6, 7)

	fmt.Println(len(slice1)) // تعداد ایتمی که الان داره
	fmt.Println(cap(slice1)) // تعداد ایتمی که میتونه درحال حاظر بپذیره
	// -----------مقدار الان داره ===> 5
	// -----------مقدار ظرفیت الان داره ===> 10
	slice2 := make([]int, 5, 10) // تعداد 5 ایتمی که الان داره و 10 تعداد ایتمی که میتونه درحال حاظر بپذیره

	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	fmt.Println(slice2)

	fmt.Println(slice2[2])

	// var slice3 []int  البته با این شکل هم خودش  یک ضرفیت در نظر میگیره که بیشتر از مقدار که داره هست

	slice4 := slice1[2:5] // قسمت 2 رو تا 5 جدا میکنه و داخل میریزه البته بجر خود 5

	fmt.Println(slice4)

}
