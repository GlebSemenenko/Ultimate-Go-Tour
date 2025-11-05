package main

import "fmt"

func main() {
	const ui = 123
	const uf float32 = 3.141592
	// const myUint8 uint8 = 1000 // переполнили - ошибка компиляции

	const ( // блок констант
		Pi       = 3.14
		Language = "Go"
		IsActive = true
	)

	const ( // можно с указанием типа
		a int     = 10
		b float64 = 20.5
		c string  = "Hello"
	)

	const ( // блок констант
		A2 = iota // 0 : Начинается с 0
		B2        // 1 : Увеличивается на 1
		C2        // 2 : Увеличивается на 1
	)

	fmt.Println(A2, B2, C2)
}