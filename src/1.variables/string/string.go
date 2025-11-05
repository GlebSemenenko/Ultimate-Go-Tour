package main

import "fmt"

func main() {

	str1 := "1hello"
	str2 := "Фhello"

	// str3 := '' // ординарные кавычки не работают

	strWitOutFormat := '\n' // не может содержать спец символы
	fmt.Println(strWitOutFormat)

	multistring := `
	1
	2
	3` // многострочная строка

	fmt.Println(multistring)

	// len возвращает кол-во байт в строке
	fmt.Println(len(str1)) // 6
	fmt.Println(len(str2)) // 7 тк Ф весит в UTF-8 2 байта

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
