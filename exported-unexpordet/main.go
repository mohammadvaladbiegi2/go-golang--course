package main

//  برای این که بتونی قابلیت export داشته باشی باید go.mod به پروژه اضافه کنی با دستور زیر

//  go mod init <module-name>

import (
	"exported-unexpordet/student"
	"fmt"
)

func main() {

	//  func / struct / varibale / const اگراول اسمشون بزرگ باشه در بقیه فایل ها در دسترس هست اما کوچیک فقط داخل اون پکیج

	user1 := student.Student{
		Name:  "mamad",
		ID:    1345645,
		Score: 15.60,
	}

	fmt.Println(user1)
	fmt.Println(student.Namepublic)
	fmt.Println(student.Funcpublic())
}
