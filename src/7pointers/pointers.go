package main

import (
	"fmt"
	"reflect"
)

// * указатели - это типы данных которые в качестве значения (по сути ссылки)
// хранят адресс ячейки в памяти либо другой указатель (может быть nil)

func main () {
	var pointer *int
	fmt.Println(pointer) // ничего не указали = nill

	var a int = 55
	fmt.Println(a) // ничего не указали = nill

	var poiterA = &a // создаем указатель на переменную
	fmt.Printf("%T - %#v - %#v - ", poiterA, poiterA, *poiterA) // распечатает кэш как в джаве
	fmt.Println(reflect.TypeOf(poiterA)) // смотрим тип данных переменной

	var newPoint = new(float64) // если создаем так то можно менять значене переменной которая лежит по указателю
	*newPoint = 10.555
	fmt.Println(*newPoint)

	c := 10 
	ptr := &c

	t1(ptr)

	cd := 900
	ptr2 := &cd

	t2(ptr, ptr2)
	fmt.Println(c, cd, " after swap")




}	

// Напишите функцию, которая принимает указатель на целое число и увеличивает его значение на 10.
func t1(c *int) { // так меняем переменную глобально
	*c += 10
}

func t2(point *int, point2 *int) {
	a := *point 
	*point = *point2
	*point2 = a 
}

// Напишите функцию, которая создает новый массив из 5 элементов, заполненных нулями, и возвращает указатель на него.
func t3 () *[5]int  {
	var data [5] int
	datapoint := &data
	return datapoint
}

// Создайте структуру Person с полями Name и Age. Напишите функцию, которая принимает указатель на Person и увеличивает возраст на 1.
type Person struct {
	Name string
	Age int  
}

func t4 (p *Person) {
	p.Age ++
}

// указатели можно использовать для экономии памяти
// если передавать просто структуру в метод, то она будет корироваться -> слайсы позволяют этого избежать