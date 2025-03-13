package main

import "fmt"

// main نقطه ورود برنامه است
func main() {

	// یک حلقه for که اعداد زوج از 0 تا 10 را چاپ می‌کند
	for i := 0; i < 10; i += 2 {
		fmt.Println(i)
	}

	// یک حلقه for که اعداد را از 0 تا کمتر از 100 با گام 10 چاپ می‌کند
	j := 0
	for j < 100 {
		fmt.Println(j)
		j += 10
	}

}
