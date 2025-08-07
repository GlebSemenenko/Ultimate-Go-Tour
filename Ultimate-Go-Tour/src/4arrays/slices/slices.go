package main

import (
	"fmt"
)

func main() {
	arr := [10]int{1,2,3}

	slice := arr[1: 9] // с индекса 2 по 8
	// slice := arr[:] // копируем массив в слайс
	// slice := arr[:5] // с 0 по 4 
	// slice := arr[2:] // c 3 по конец
	// sliceFromSlice := slice[:] // создание слайса из другого слайса
	// s := make([]int, 3, 5) // Создает слайс длиной 3 и емкостью 5
	// var s []int // пустой слайс 0 длинны и емкости // или так s := []int{} 
	fmt.Println(slice, "Создание слайса из части массива")

	slice[1] = 19000 // изменяем значение по индексу
	slice = append(slice, 1000) // добавляем элемент
	slice = append(slice, 1,2,33,4,5)
	slice = append(slice, slice...) // если нужно добавить слайс в  слайс
	fmt.Println(slice)

	tz()
} 

func capAndLen () {
	var arr []int 
	fmt.Println(cap(arr), " -cap ",len(arr)," -len")
	for i:=0; i < 10; i++ {
		arr = append(arr,i)
		fmt.Println(cap(arr), " -cap ",len(arr)," -len")
	}
}

func tz () {
	var arr []int
	for {
		var a int 
		fmt.Scan(&a)
		if (a == 010) {
			fmt.Println("конец программы")
			break
		}
		arr = append(arr, a)
		fmt.Println(arr)
	}
}