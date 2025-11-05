package main

import (
	"fmt"
	"math"
)

func main() {
	a := 10
	b := 3

	a %= b
	fmt.Println(a)

	fmt.Println(a+b, "сложение")           // 13
	fmt.Println(a-b, "вычитание")          // 7
	fmt.Println(a*b, "умножение")          // 30
	fmt.Println(a/b, "деление")            // 3
	fmt.Println(a%b, "остаток от деления") // 1

	a += b
	a -= b
	a *= b
	a /= b
	a %= b
	a++
	a--

	c := 0
	for i := 0; i < 5; i++ {
		c++
		fmt.Println(c)
	}

	fmt.Println(math.Pow(2, 4)) // в функции возведения в степень используется тип данных float
	var sla rune
	fmt.Println(sla)
}
