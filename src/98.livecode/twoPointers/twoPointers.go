package main

import (
	"fmt"
	"strings"
)

func main () {

}

func isPalindrome(s string) bool {
    slower := strings.ToLower(s)
	trim := strings.TrimSpace(slower)
	nospaces := strings.ReplaceAll(trim, " ","")

	arr := []rune(nospaces)	
	r := len(arr) -1
	for l := 0; l < r; l++ {
		if (arr[l] != arr[r]) {
			return false
		}
		r--
	}
	return true
}


