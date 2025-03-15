package main

import (
	"fmt"
	"time"
)

func main() {
	var hour, minute, second int
	for {
		second++
		if second >= 60 {
			second = 0
			minute++
		}
		if minute >= 60 {
			hour++
			minute = 0
		}
		if hour >= 24 {
			hour = 0
		}
		fmt.Printf("%02d:%02d:%02d\n", hour, minute, second)
		time.Sleep(time.Second)
	}
}
