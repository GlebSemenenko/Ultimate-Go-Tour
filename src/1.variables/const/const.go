package main

func main() {
	const ui = 123
	const uf float32 = 3.141592  
	// const myUint8 uint8 = 1000 // переполнили - ошибка компиляции


	const ( // блок констант
    A2 = iota  // 0 : Начинается с 0
    B2         // 1 : Увеличивается на 1
    C2         // 2 : Увеличивается на 1
)
}