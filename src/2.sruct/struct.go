package main

import "fmt"

// структуры данных - это аналоги классов

func main() {

	var map1 = map[string]int{
		"23r2r2Key_1": 1,
		"erteey_2":    2,
		"Kteey_3":     3,
		"erteKey_4":   4,
		"Keyegerge_5": 5,
		"egKey_6":     6,
	}

	for key, value := range map1 {
		fmt.Println(key, value)
	}

	fmt.Println(map1)
}

type example struct {
	flag    bool
	counter int16
}
