package main

import (
	"fmt"
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
	fmt.Println(isAnagram("asd", "das"))
	fmt.Println(isAnagram("qwert", "qwerty"))
	fmt.Println(isAnagram("qwe", "qwe"))
	fmt.Println(isAnagram("ahgf", "afgh"))
	fmt.Println(isAnagram("poiuy", "yuiop"))
	fmt.Println(isAnagram("qwedas", "cxz"))
}
