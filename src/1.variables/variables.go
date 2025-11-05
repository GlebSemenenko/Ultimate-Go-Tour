package main

import (
	"fmt"
)

var a = "global" // Объявление переменной на уровне пакета
// aa := 111 глобальные переменные нельзя объявлять коротким синтаксисом

const name = "const" // Не меняют значение

func main() {

	// ПЕРЕМЕНННЫЕ В GO

	a := 10
	_ = a
	var rsr = 10
	_ = rsr

	var i int8 = 127
	_ = i // если объявил переменную, то ее обязательно использовать, но можно использовать пустой идентификатор чтобы это обойти
	var i1 int16 = 2
	var i2 int32 = 2
	var i3 int64 = 2
	fmt.Print(i1, " ", i2, " ", i3, "\n") // функция Print может принимать сколько угодно аргументов

	var f32 float32 = 12.2
	var f64 float64 = 12.5
	fmt.Print(f32, " ", f64, "\n")

	var s string = "str"
	fmt.Print(s, "\n")

	var name string = "local" // локальная переменная перебьет глобальную
	fmt.Print(name, "\n")

	var bool = true
	fmt.Print(bool)

	// ТИПОБЕЗОПАСНОСТЬ И ПРИВЕДЕНИЕ ТИПОВ

	// fmt.Println(i+i3) так нельзя в go приведение типов всегда явное
	fmt.Print(i+int8(i3), "\n") // В даже при переходе в разные int нужно прописывать руками

	var binary int = 0b10101101
	fmt.Println(binary) // 173

}
