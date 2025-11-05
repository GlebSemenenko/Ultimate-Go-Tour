package main

import (
	"fmt"
	"sort"
	"strconv"
)


func main () {
	str := "йцу"
	str2 := "wee"
	fmt.Println(str == str2)


	runeArr1 := []rune(str)
	runeArr2 := []rune(str2)
	
	v := 190
	vstr := string(v)
	fmt.Println(vstr)

	s := strconv.Itoa(v)
	fmt.Println(s)


	fmt.Println(len(runeArr1), len(runeArr2))

	arr := [...]int {1,2,3}
	last := arr[len(arr)-1]
	fmt.Println(last,"последний элемен")
	arr2 := [...]int {1,2,13}
	fmt.Println(arr == arr2)


}


func heightChecker(arr []int) int { // 1051
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(arr)
	res := 0
	for i, v := range sorted {
		if arr[i] != v {
			res ++
		}
	}
	return res
}

func findLUSlength(a, b string) int { // 521
	if(a == b) {
		return -1
	} else {
		runesA := []rune(a)
		lenAString := len(runesA)
		
		runesB := []rune(b)
		lenBString := len(runesB)
		if (lenAString > lenBString) {
			return lenAString 
		} else {
			return lenBString
		}
	}	
}

// func countLetters(input string) int{
// 	res := 0

// }