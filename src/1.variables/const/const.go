package main

import "fmt"

func main() {
	const ui = 123
	const uf float32 = 3.141592
	// const myUint8 uint8 = 1000 // переполнили - ошибка компиляции

	const ( // блок констант
		Pi              = 3.14
		Language string = "Go"
		IsActive        = true
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

	// Пример 1: пропуск значения iota с помощью _
	const (
		_     = iota // 0 - пропускаем
		One          // 1
		Two          // 2
		Three        // 3
	)

	// Пример 2: битовые маски (флаги) через сдвиг 1 << iota
	const (
		FlagRead  = 1 << iota // 1 << 0 = 1
		FlagWrite             // 1 << 1 = 2
		FlagExec              // 1 << 2 = 4
	)

	// Пример 3: арифметические преобразования iota
	const (
		ValFirst  = iota * 10 // 0 * 10 = 0
		ValSecond             // 1 * 10 = 10
		ValThird              // 2 * 10 = 20
	)

	fmt.Println("A2, B2, C2 =", A2, B2, C2)
	fmt.Println("One, Two, Three =", One, Two, Three)
	fmt.Println("Flags (Read, Write, Exec) =", FlagRead, FlagWrite, FlagExec)
	fmt.Println("ValFirst, ValSecond, ValThird =", ValFirst, ValSecond, ValThird)
}
