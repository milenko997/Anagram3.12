package main

import (
	"bufio"
	"fmt"
	"os"
)

func isAnagram(a, b string) bool {
	// make maps that will count how many times each rune is repeated in string
	count1 := make(map[rune]int)
	count2 := make(map[rune]int)
	// compare if they are the same length
	if len(a) != len(b) {
		return false
	}
	// counting each rune from each string
	for _, c := range a {
		count1[c]++
	}
	for _, c := range b {
		count2[c]++
	}
	// compare first string with second
	for first, second := range count1 {
		if count2[first] != second {
			return false
		}
	}
	// compare second string with first
	for first, second := range count2 {
		if count1[first] != second {
			return false
		}
	}
	return true
}

func main() {
	input1 := bufio.NewScanner(os.Stdin)
	input2 := bufio.NewScanner(os.Stdin)
	fmt.Println("Unesite prvu rec: ")
	input1.Scan()
	fmt.Println("Unesite drugu rec: ")
	input2.Scan()

	fmt.Println(isAnagram(input1.Text(), input2.Text()))
}
