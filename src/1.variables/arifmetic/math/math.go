package main

import (
	"fmt"
	"math"
)

// все функции работают с float64

func main() {

	fmt.Println(math.Pi) // pi
	i := math.MaxInt
	_ = i

	fmt.Println(math.Sqrt(4))   // корень
	fmt.Println(math.Abs(-4))   // модуль
	fmt.Println(math.Log(28))   // логарифм
	fmt.Println(math.Pow(2, 2)) // степень

	// Округление
	fmt.Println(math.Round(1.4)) // до ближайшего целого
	fmt.Println(math.Floor(1.9)) // округление вниз
	fmt.Println(math.Ceil(1.9))  // округление вверх
	fmt.Println(math.Trunc(1.1)) // удаление дробной части

}
