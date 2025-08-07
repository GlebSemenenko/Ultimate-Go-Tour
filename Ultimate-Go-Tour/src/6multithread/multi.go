package main

import (
	"fmt"
	"sync"
)

func main () {
	zeroToMany(10000)
	fmt.Println("end")
}

func withWait() {
	var wg sync.WaitGroup
	
	for i := 0; i < 10 ; i++ {
		wg.Add(1) // добавлять потоки нужно изнутри горутины!
		go func (i int)  {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait() // блокируем горутину она разблочится когда в вейт группе останется 0 задач
}

/* Алгоритм работы с WaitGroup
var wg sync.WaitGroup создали вейт группу
wg.Add добавили горутину
wg.Done вызываем сразу после выполнения задачи
*/ 


func zeroToMany (count int) {
	r := 0
	var wg sync.WaitGroup
	var mutex sync.RWMutex // использование мьютекса исключает ситуацию когда нескуолько горутин одновременно читают или изменяют значение переменной
	wg.Add(count)

	for i := 0; i < count; i++ {
		
		go func() {
			defer wg.Done()
			mutex.Lock()
			r++
			mutex.Unlock() 
		}()	
	}
	wg.Wait()
	fmt.Println(r)
}