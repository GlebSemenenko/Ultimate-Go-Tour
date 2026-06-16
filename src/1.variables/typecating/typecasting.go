package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	i := 2
	f := 3.14
	s := "3"

	// Float -> Int
	fmt.Println(f+float64(i), " f + i") // Разрядность важна !
	fmt.Println(int(f)+i, " f + i")

	// String -> Int
	intFromStr, _ := strconv.Atoi(s) // строка -> int
	fmt.Println(intFromStr + i)

	// Int -> String
	strFromInt := strconv.Itoa(42)
	fmt.Println(strFromInt)

	// Проверяем тип данных в го с помощью рефлексии
	ref := reflect.TypeOf(f)
	fmt.Println(ref)

	nnn := -1
	n := uint(nnn)
	fmt.Println(n)     // 18446744073709551615'
	fmt.Println(n + 1) // 0
}
