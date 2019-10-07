package main

import "fmt"

func main() {
	result := isStringPalidrome("mom")
	var wordResult string

	if result == true {
		wordResult = "is"
	} else {
		wordResult = "is not"
	}

	fmt.Printf("Your word %s a palindorome", wordResult)
}

func isStringPalidrome(str string) bool {
	rs := []rune(str)
	lenOfRs := len(rs)

	for i, j := 0, lenOfRs-1; i < lenOfRs/2; i, j = i+1, j-1 {
		if rs[i] != rs[j] {
			return false
		}
	}
	return true
}

// longest palindrome in substr
