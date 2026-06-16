package main

import "fmt"

func main() {
	
	s := "hello, joe !"
	_ = s

	runes := []rune(s)
	runes1 := []rune(s)

	// for range в значении возвращаю ссылку. Будь внимателен!
	for _, v := range runes {
		if v == 'h' {
			v = 'r' // не поменяется так как это копия
		}
	}
	fmt.Println(string(runes1))

	for n, r := range runes1 {
		if r == 'l' {
			runes1[n] = '!' // так правильно
		}
	}
	fmt.Println(string(runes1))

	for i := 0; i < len(runes); i++ {
		if runes[i] == 'l' {
			runes[i] = '!'
		}
	}
	fmt.Println(string(runes))
}
