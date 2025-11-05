package main

import (
	"errors"
	"fmt"
)

func main() {

	var sl []int

	fmt.Println(sl == nil)

	l := [...]int{1, 3, 21, 3, 414, 5}
	listRes(l[:])
	fmt.Println(l)

}

func listRes(ls []int) {
	for i := 0; i < len(ls); i++ {
		ls[i] += 10
	}
}

func listRes1(ls []int) {
	for i := range ls {
		ls[i] += 10
	}
}

func sayHello() { // простая функция без аргументов и возвращаемых значений
	fmt.Println("Hello !")
}

// Аргументы функций

func sayName(s string) { // 1 аргумент
	fmt.Printf("Hello %s, How are you", s)
}

func sayNameAndAge(s string, i int) { // 2 аргумента
	fmt.Println("It is", s, "he is", i, "years old")
}

func sayNames(s1, s2 string) { // аргументы с 1 типом данных
	fmt.Println("Hello", s1, "and", s2)
}

func sayNames1(string) { // можно не именовать аргумент, но это неудобно
}

// Возвращаемые значения функций

func returnValue1() int {
	return 1
}

func returnValue2() (int, float32, error) {
	return 1, 2.2, nil
}

func mistakeFunc(a int) (int, error) {
	if a == 0 {
		return 0, errors.New("на ноль делить нельзя")
	}
	return 1, nil
}

// Defer - отложенный вызов

func deferFunc() { // 1 аргумент
	defer fmt.Println("End of func")
	fmt.Println("Hello")
}

func manyDeferFunc() {
	defer fmt.Println("3")
	defer fmt.Println("2")
	defer fmt.Println("1")

	fmt.Println("End")
}
