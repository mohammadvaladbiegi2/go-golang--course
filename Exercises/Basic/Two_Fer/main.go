package main

import "fmt"

func TwoFer(name string) string {
	if name == "" {
		return "One for you, one for me."
	}

	return "One for " + name + ", one for me."
}

func main() {

	fmt.Println(TwoFer(""))
	fmt.Println(TwoFer("Alice"))
	fmt.Println(TwoFer("Bob"))

}
