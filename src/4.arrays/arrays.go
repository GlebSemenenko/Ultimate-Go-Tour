package main

import (
	"fmt"
)

func main() {

	// создание массивов
	var data [5]int
	fmt.Println(data) // zero val initialization

	slice := []int{1, 2, 3, 4, 5}
	arr := [5]int(slice)
	fmt.Println(arr)

	var dataM [2][3]int
	fmt.Println(dataM) // многомерный массив

	data2 := [...]int{1, 2, 3, 4, 5} // [1 2 3 4 5]
	fmt.Println(data2)

	data3 := [5]int{1, 2, 3}
	fmt.Println(data3)

	data4 := [5]int{3: 4}
	fmt.Println(data4)

	data5 := [5]int{2: 5, 6, 1: 7}
	fmt.Println(data5)

	data6 := [4]int{1, 2, 3}
	fmt.Println(data6, "data6")
	data7 := [4]int{}
	data6 = data7

	var emptyArr []int
	l := []int{}

	fmt.Println(emptyArr == nil)
	fmt.Println(l == nil)

	fmt.Println(data == data2)
	// fmt.Println(data >= data2) сравнивать можно только больше или равно
	// make и append на массив использовать нельзя

	// массивы не имеют постоянного места в памяти

	fmt.Println("_________________")
	for _, v := range data3 {
		v *= 100
	} // если меняем значение через range оно копируется и сам массив остается неизменным
	fmt.Print(data3)

	for i := 0; i < len(data); i++ {
		data[i] += 100
	}
	fmt.Println(data)

	processArray(data[:])
	fmt.Println(data, "вызов processArray")

	fmt.Println("-------------------")

	sl := []int{1, 2, 3}
	appendInSlice(sl[:1])
	fmt.Println(sl)

	// АЛЛОКАЦИЯ МАССИВОВ

}

// При передаче в функцию массив копируется
func processArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] += 5555500000
	}
}

func appendInSlice(arr []int) {
	_ = append(arr, 19999999)
}

func withpointer(arr *[3]int) { // значение массива не будет изменено
	for i, v := range *arr {
		arr[i] = v * 2
	}
}

func withnopointer(arr *[3]int) { // значение массива будет изменено
	for i, v := range arr {
		arr[i] = v * 2
	}
}

func withSlice(arr []int) { // значение слайса будет изменено
	for i, v := range arr {
		arr[i] = v * 2
	}
}

func withoutMutation(original []int) []int { // значение слайса не будет изменено
	result := make([]int, len(original))
	for i, v := range original {
		result[i] = v * 2
	}
	return result
}
