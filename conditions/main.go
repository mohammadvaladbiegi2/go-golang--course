package main

import "fmt"

// main نقطه ورود برنامه است
func main() {

	// تعریف و مقداردهی اولیه متغیر age
	age := 18

	// بررسی شرطی برای تعیین امکان رأی دادن بر اساس سن
	if age > 18 {
		fmt.Println("you can vote")
	} else if age == 18 {
		fmt.Println("you can vote but you need to come with parent")
	} else {
		fmt.Println("you can't vote")
	}

	// تعریف و مقداردهی اولیه متغیر score
	score := 80

	// استفاده از switch برای ارزیابی نمره و چاپ پیام مناسب
	switch score {
	case 100:
		fmt.Println("you are perfect")
	case 80:
		fmt.Println("you are good")
	case 60:
		fmt.Println("you are pass")
	default:
		fmt.Println("you are fail")
	}

}
