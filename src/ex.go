package main

import "fmt"

func main() {

	// Создать мап 4 способа
	map1 := make(map[string]int)
	map2 := make(map[string]int, 20)

	var map3 = map[string]int{
		"A": 1,
		"B": 2,
	}

	var map4 = map[string]int{}

	map1["a"] = 10
	map1["a"] = 20

	x := map1["a"]

	fmt.Println(x)

	delete(map1, "a")

	fmt.Println(map1, map2, map3, map4)
}
