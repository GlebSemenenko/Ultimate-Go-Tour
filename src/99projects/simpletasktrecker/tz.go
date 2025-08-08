package main

import (
	"fmt"
	"os"
)

func main() {
	hellostr := 
`- 1 Посмотреть закладки
- 2 Добавить закладки 
- 3 Удалить закладку
- 0 Выход`

	fmt.Println(hellostr)
	m := make(map[string]string)

	for {
		var i int
		fmt.Print("Введите номер действия: ") // Добавлено для ясности
		fmt.Scan(&i)

		switch i {
		case 0:
			fmt.Println("Bye bye")
			os.Exit(0)

		case 1:
			if len(m) == 0 {
				fmt.Println("Закладки отсутствуют") // Вывод сообщения, если закладок нет
			} else {
				fmt.Println("Закладки:", m) 
			}

		case 2:
			var k, v string
			fmt.Print("Введите название закладки: ") 
			fmt.Scan(&k)
			fmt.Print("Введите URL закладки: ") 
			fmt.Scan(&v)
			m[k] = v
			fmt.Println("Закладка добавлена:", k)

		case 3:
			var k string
			fmt.Print("Введите название закладки для удаления: ") 
			fmt.Scan(&k)
			if _, exists := m[k]; exists {
				delete(m, k)
				fmt.Println("Закладка удалена:", k)
			} else {
				fmt.Println("Закладка не найдена:", k) // Сообщение, если закладки нет
			}
		default:
			fmt.Println("Неверный ввод")
		}
	}
}
