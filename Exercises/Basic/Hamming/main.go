package main

import "fmt"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("not same length")
	}
	var result int
	for index, value := range a {
		if string(value) != string(b[index]) {
			result++
		}
	}
	return result, nil
}

func main() {

	fmt.Println(Distance("GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT"))

}
