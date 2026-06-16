package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	str1 := "1hello"
	str2 := "Фhello"

	//str3 := 'hi' // ординарные кавычки не работают это для рун
	strWitOutFormat := '\n' // не может содержать спец символы
	_ = strWitOutFormat

	multistring := `
	1
	2
	3` // многострочная строка
	_ = multistring

	// len возвращает кол-во байт в строке
	fmt.Println(len(str1))                    // 6
	fmt.Println(len(str2))                    // 7 тк Ф весит в UTF-8 2 байта
	fmt.Println(utf8.RuneCountInString(str2)) // размер строки в символах
	// строки можно сравнивать через == >= <= Cравнение будет по байтам, а не по символам

	sub := str2[:2] // так же и в слайсах -> мы получаем 2 байта, а не символа !
	fmt.Println(sub)

	for index, char := range str2 {
		fmt.Println(index, char)

	}

	runes := []rune(str2)

	fmt.Println(len(runes), "123123123")
	fmt.Println(runes)

	for i, r := range runes {
		fmt.Printf("Индекс: %d, Символ: %c\n", i, r)
	}

	r := '1' // создаем руну
	_ = r
	fmt.Println(r) // выводит номер из UTF-8
}

// Строка - Immutable sequences of bytes representing UTF-8 encoded text.
// Создание строки
// Длина строки в байтах и символах
// Подстроки
// Сравнение строк
