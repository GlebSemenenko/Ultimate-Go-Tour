package main

import (
	"fmt"
)

func main () {
	data := [...] int {2,2,3,2,5} 
	
	// sl = data[:]

	fmt.Println(uniqueArr(data[:]))

}

// сумма  элементов массива
func sumOfArray (inputArr []int) int {
	res := 0
	for _, v := range inputArr {
		res += v
	}
	return res
}

// найти уникальные значения в массиве
func uniqueArr (inputArr []int) []int {

	res := make([]int, 0, 0)

	m := make(map[int]string)
	for _, v := range inputArr {
		m[v] = ""
	}
	fmt.Println(m)
	for k, _ := range m {
		res = append(res, k)
	}
	return res
}