package main

import "fmt"

// main نقطه ورود برنامه است
func main() {

	// تعریف یک ساختار به نام User با فیلدهای مختلف
	type User struct {
		name     string
		age      int
		email    string
		password string
		isActive bool
	}

	// ایجاد یک نمونه از ساختار User با استفاده از مقداردهی اولیه
	user1 := User{
		name:     "ali",
		age:      20,
		email:    "ali@gmail.com",
		password: "123456",
		isActive: true,
	}

	// چاپ اطلاعات کاربر user1
	fmt.Println(user1)

	// ایجاد یک نمونه دیگر از ساختار User و مقداردهی فیلدها به صورت جداگانه
	var user2 User
	user2.name = "reza"
	user2.age = 21
	user2.email = "reza@gmail.com"
	user2.password = "123456"
	user2.isActive = true

	// چاپ اطلاعات کاربر user2
	fmt.Println(user2)

	// ایجاد و مقداردهی یک نمونه دیگر از ساختار User به صورت مستقیم
	user3 := User{"ali", 20, "ali@gmail.com", "123456", true}
	// چاپ اطلاعات کاربر user3
	fmt.Println(user3)

}
