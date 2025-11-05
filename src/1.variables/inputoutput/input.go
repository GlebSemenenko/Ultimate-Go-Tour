package main

import (
	"fmt"
)

func main() {
	name := "hello"
	_, _ = fmt.Scan(&name)
	fmt.Println(name)
}
