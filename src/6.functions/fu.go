package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	funcWithManiArguments(1, 3, 3, 4, 3, 43, 4, 4, 2, 4, 4, 45)
}

func simple() {
	println("simple")
}

func simpleWithArgs(i int) {
	fmt.Println("simpleWithArgs ", i)
}

func funcWithReturnValue(i int) int {
	return i * i
}

func funcWithReturnValues(s string) (int, error) {
	if s == "" {
		return 0, errors.New("s is empty") // Важно писать о конкретной ошибке.
	}
	return strconv.Atoi(s)
}

func funcWithNamedReturnValues() (n int, err error) {
	err = nil
	n = 1
	return n, err
}

func funcWithManiArguments(i ...int) {
	fmt.Println(reflect.TypeOf(i)) // [] int
	for c, v := range i {
		fmt.Println(c, v)
	}
}

type Example struct {
	a int
	b string
}

// В Golang есть различия между методами и функциями. Функция свободна, а метод привязан к какой-либо структуре.
// Метод с получателем-указателем (можем менять структуру)
func (r *Example) area() (int, string) {
	return r.a, r.b
}

// Метод с получателем-значением (копирует структуру)
func (r Example) perim() (int, string) {
	return 2 * r.a, r.b
}
