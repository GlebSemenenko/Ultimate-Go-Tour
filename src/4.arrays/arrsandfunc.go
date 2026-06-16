package main

import "fmt"

func main() {

	arr := [3]int{1, 2, 3}
	sl := []int{1, 2, 3}
	withnopointer(&arr)
	fmt.Println(arr)

	withSlice(sl)
	fmt.Println(sl)

}

//func withpointer(arr *[3]int) { // значение массива не будет изменено
//	for i, v := range *arr {
//		arr[i] = v * 2
//	}
//}
//
//func withnopointer(arr *[3]int) { // значение массива будет изменено
//	for i, v := range arr {
//		arr[i] = v * 2
//	}
//}
//
//func withSlice(arr []int) { // значение слайса будет изменено
//	for i, v := range arr {
//		arr[i] = v * 2
//	}
//}
//
//func withoutMutation(original []int) []int { // значение слайса не будет изменено
//	result := make([]int, len(original))
//	for i, v := range original {
//		result[i] = v * 2
//	}
//	return result
//}
