package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Конструкторы map
	m1 := map[string]string{ // создание через литерал
		"Key_1": "Value_1",
	}

	m2 := make(map[string]int)     // базовое создание
	m3 := make(map[string]int, 10) // initial capacity 10
	var m4 map[string]int          // пустая карта

	fmt.Println(m1, m2, m3, m4)

	// API map
	m1["Key_2"] = "Val_2"   // добавить элемент
	m1["Key_2"] = "Value_2" // изменить

	delete(m1, "Key_2") // удаление элемента
	delete(m1, "Key_2") // если удалить элемент, которого нет, то все ок

	m1["Key_2"] = "Value_2" // получение элемента по ключу

	// Ключом в мап могут быть только Comparable типы (типы которые можно сравнить через ==)
	// map5 := make(map[[]string]string) // так нельзя
	// Например: int, int32, float32, float64

	// проверяем наличие ключа в map
	_, ok := m1["Key_2"] // i - это значение ключа ok - булево значение, сравнение с i оно true/false
	// fmt.Println(i, "вывод значения")
	fmt.Println(ok, "вывод результата")

	/*
	   	// Итерация по map

	   	for i := 0; i < 1000; i++ { // Заполняем карту 1000 парами
	   		strK := "Key_" + fmt.Sprint(i)
	   		strV := "Value_" + fmt.Sprint(i)
	   		m1[strK] = strV
	   	}
	   //	fmt.Println(m1) // ЭЛЕМЕНТЫ В Map выводятся в случайном порядке !

	   	for k, v := range m1 { // итерация по ключам и значениям
	   		fmt.Println(k,v)
	   	}

	   	for k := range m1 { // итерация по ключам
	   		fmt.Println(k)
	   	}

	   	for _, v := range m1 { // итерация по значениям
	   		fmt.Println(v)
	   	}
	*/

	// Сравнение мап
	fmt.Println(reflect.DeepEqual(m1, m2))

}
