package main

import (
	"sort"
	"fmt"
)

func main () {

	// Сортировка списка int 
	arr := []int{1,123,1,3,131,31,1}
	sort.Ints(arr)
	fmt.Println(arr)

	// Проверяем отсортирован ли массив
	sort.SliceIsSorted(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	// Сортировка строк
	strings := []string{"б","а","г","44","ф"} // сортирует по алфамиту, алфавит определяет UTF-8
	sort.Strings(strings)
	fmt.Println(strings)

	
	fmt.Println(sort.StringsAreSorted(strings))
}