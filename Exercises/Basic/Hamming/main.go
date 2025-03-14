package main

import "fmt"

func Hamming(dna1, dna2 string) []string {

	var result []string
	if len(dna1) == 0 && len(dna1) == 0 {
		return result
	}

	for index, value := range dna1 {
		if string(value) != string(dna2[index]) {
			result = append(result, string(dna2[index]))
		}
	}
	return result
}

func main() {

	fmt.Println(Hamming("GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT"))

}
