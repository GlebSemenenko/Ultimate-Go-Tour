// программа считает индек массы тела

package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main () {
	var height float32
	var weight float32

	fmt.Println("weight")
	fmt.Fscan(os.Stdin, &weight)

	fmt.Println("height")
	fmt.Fscan(os.Stdin, &height)
	fmt.Println(imt(height,weight))
}

func imt (height, weight float32) (float32, error) {
	if height <= 0 || weight <= 0 {
		return 0, errors.New("некорректные значения: рост и вес должны быть положительными")
	}
	res := weight / (height*height)
	return float32(roundToN(float64(res),4)), nil
}

func roundToN(x float64, n int) float64 {
    factor := math.Pow(10, float64(n))
    return math.Round(x*factor) / factor
}
