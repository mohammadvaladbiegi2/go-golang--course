package main

import "fmt"

type User struct {
	ID   int
	name string
}

func main() {
	user1 := new(User)

	fmt.Println(user1)

	user2 := &User{}

	fmt.Println(user2)
}
