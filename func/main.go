package main

import "fmt"

// main نقطه ورود برنامه است
func main() {
	// نتیجه تابع sayHi را چاپ می‌کند
	fmt.Println(sayHi("mamad", 20))
	// نتیجه تابع add را چاپ می‌کند
	fmt.Println(add(1, 2))
	// نتیجه تابع power را چاپ می‌کند
	fmt.Println(power(2, 3))
}

// تابع sayHi یک پیام خوش‌آمدگویی و سن را برمی‌گرداند
func sayHi(name string, age int) (string, int) {
	return "hi " + name, age
}

// تابع add مجموع دو عدد صحیح را برمی‌گرداند
func add(a int, b int) int {
	return a + b
}

// تابع power حاصل‌ضرب دو عدد صحیح را برمی‌گرداند
func power(a int, b int) (result int) {
	result = a * b
	return
}
