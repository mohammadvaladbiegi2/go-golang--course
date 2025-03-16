package main

import "fmt"

func main() {

	//  [] اول تایپ کلید مشخض میکنم و  در ادامه تایپ ولیو رو
	user := make(map[string]string)
	user["name"] = "mamad"
	user["email"] = "mohammad@gmail.com"
	user["password"] = "1234556"

	fmt.Println(user)

	mapValue, haveValue := user["name"]

	fmt.Println("map Value", mapValue, "\n", " ====>", haveValue)

	//  حذف یک مقدار داخل مپ

	delete(user, "email")

	fmt.Println("user ==> ", user)

}
