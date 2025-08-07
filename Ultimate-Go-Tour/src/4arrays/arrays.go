package main

import (
	"fmt"
)

func main () {
	// создание массивов
	var data [5] int 
	fmt.Println(data) // zero val initialization

	var dataM [2][3]int 
	fmt.Println(dataM) // многомерный массив 

	data2 := [...] int {1,2,3,4,5} // [1 2 3 4 5]
	fmt.Println(data2)

	data3 := [5]int{1,2,3}
	fmt.Println(data3)

	data4 := [5]int {3:4}
	fmt.Println(data4)

	data5 := [5]int{2:5, 6, 1:7}
	fmt.Println(data5)

	fmt.Println(data == data2)
	// fmt.Println(data >= data2) cравнивать можно только болоьше или равно
	// make и apend на массив использовать нельзя

	// массивы не имеют постоянного места в памяти

	fmt.Println("_________________")
	for _, v := range data3 {
		v *=100
	} // если меняем значение через range оно копируется и сам массив остается неизменным
	fmt.Print(data3)

	for i := 0; i < len(data); i++{
		data[i] += 100
	}
	fmt.Println(data)

	processArray(data[:])
	fmt.Println(data, "вызов processArray")
	
	// АЛЛОКАЦИЯ МАССИВОВ

}

// При передачи в функцию макссив копируется
func processArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] += 5555500000
	}
}