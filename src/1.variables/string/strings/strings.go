package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Библиотека для работы со строками го
// Одновременно является Стринг билдером
func main () {

	str := "start"

	// Клонирование строки. Жрет память !
	clonestr := strings.Clone(str)
	fmt.Println(clonestr, "клонировал строку")

	// Сравнение строк. 
	// Лучше сравнивать по ==
	// Возвращает int: -1, 0, 1
	fmt.Println(strings.Compare(str, clonestr))

	//Проверяем наличие подстроки
	sub := "rt"
	fmt.Println(strings.Contains(str,sub))

	// Проверяем наличие руны
	subRune := 't'
	fmt.Println(strings.ContainsRune(str, subRune))

	//  Считаем сколько раз 
	countStr := "aabbccddeeff"
	fmt.Println(strings.Count(countStr,"a"))

	// Верхний и нижний регистр
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))


	fmt.Println(strconv.Atoi("1200")) // из строки в число
	fmt.Println(strconv.Itoa(12)) // из числа в строку
	//fmt.Println(int("1200"))

}


