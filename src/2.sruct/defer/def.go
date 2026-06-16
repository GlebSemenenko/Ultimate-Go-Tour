package main

import "fmt"

// defer - отложенный вызов функции
// помещает функцию в стек (это очень важно) отложенных вызовов
// используется для закрытия ресурсов, например конект к БД

func main() {
	a()
	a1()
	a2()

	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
}

func d() {
	defer fmt.Println("defer from d")
	fmt.Println("hello from d")
}

func d1() {
	fmt.Println("hello from d1")
	defer fmt.Println("defer from d1")
}

func d2() {
	fmt.Println("defer from d2")
	panic("panic from d2")
}

func a() {
	defer fmt.Println("a")
	fmt.Println("a")
}
func a1() {
	defer fmt.Println("a1")
	fmt.Println("a1")
}
func a2() {

	defer fmt.Println("a2")
	fmt.Println("a2")
}
