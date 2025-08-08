package main

import "fmt"

// структуры данных - это аналоги классов

func main () {
	Greet()
	PersonGreet("Joe")
	PersonGreetAndEge("Sam", 14)
	a,_,_ := Bezum() // так доставать значения если их несколько
	fmt.Println(a)
}

// без аргументов
func Greet() {
	fmt.Println("Hello")
}
// возвращаемые значения
// в го нет перегрузки и переопределения методов
func PersonGreet(name string) string {
	fmt.Println("Zdarov ", name)
	return "str"
}

func PersonGreetAndEge(name string, age int) {
	fmt.Println("Zdarov ", name)
}

// можно возвращать несколько значений, обычно это делается чтобы обрабатывать ошибки
func Tarkov() (int, error) {
	return 1488, nil
}

// если возвращаемых значений несколько лучше сделать их именнованными 
func Bezum() (a int, b int, c int) {
	return 1, 2, 3
}

// func Dota(s PersonGreet() ) int {
// 	PersonGreet("wewe")
// 	return 1 
// }

type example struct {
	flag bool
	counter int16
}