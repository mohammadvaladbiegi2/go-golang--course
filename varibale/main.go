package main

import "fmt"

// main نقطه ورود برنامه است
func main() {

	// تعریف یک متغیر رشته‌ای به نام name
	var name string

	// تعریف دو متغیر رشته‌ای به نام‌های email و password
	var email, password string

	// تعریف و مقداردهی اولیه یک متغیر عدد صحیح به نام age
	age := 21

	// تعریف و مقداردهی اولیه سه متغیر num1، num2 و char
	num1, num2, char := 1, 2, 'a'

	// تعریف و مقداردهی اولیه یک متغیر عدد اعشاری به نام numfloat
	numfloat := 6.13

	// تبدیل عدد اعشاری به عدد صحیح و ذخیره در numint
	numint := int(numfloat)

	// تعریف یک ثابت به نام pi
	const pi = 3.14

	// چاپ مقادیر متغیرها و ثابت‌ها
	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(password)
	fmt.Println(age)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(char)
	fmt.Println(numint)
	fmt.Println(pi)
}
