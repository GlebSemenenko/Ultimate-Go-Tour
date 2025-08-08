package main

import (
	"fmt"
)

func main() {

	// СОЗДАНИЕ МАПЫ 

	cityMap := map[string]string{
		"Moscow": "RU",
		"London": "EN",
		"Tarkov": "RU",
		"LA": "US",
	}

	m := make(map[string]int)

	fmt.Println(cityMap, m) // вывод в случайном порядке !

	cityMap["Dublin"] = "FR" // ДОБАВЛЕНИЕ ЭЛЕМЕНТОВ 
	cityMap["Dublin"] = "IR" // ИЗМЕНЕНИЕ ТАКЖЕ
	delete(cityMap, "Dublin") // УДАЛЕНИЕ
	val, exist := cityMap["Dublin"] // ПРОВЕРКА НАЛИЧИЯ КЛЮЧА
	fmt.Println(val, exist)

	// иттерация по мап 
	for k, v := range cityMap { // вывод в случайном порядке !
		fmt.Println(k,v)
	}
	
}