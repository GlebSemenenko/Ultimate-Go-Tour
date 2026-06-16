package main

import "fmt"

func main() {

	var pot int = 20
	fmt.Println(&pot)

	var potpointer = new(int16)
	fmt.Println(*potpointer)

}
