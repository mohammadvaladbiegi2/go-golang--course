package main

import (
	"fmt"
	"log"
)

func main() {

	var nameServer string
	var codeServer int
	// نمایش پیام راهنما و خواندن نام سرور
	fmt.Print("Enter the server Name: ")
	fmt.Scanln(&nameServer)

	// نمایش پیام راهنما و خواندن کد سرور
	fmt.Print("Enter the server Code: ")
	fmt.Scanln(&codeServer)

	if nameServer == "" || codeServer == 0 {
		log.Fatalf("name server is empty or not set server code")
	}
	fmt.Println(nameServer)
	fmt.Println(codeServer)

}
