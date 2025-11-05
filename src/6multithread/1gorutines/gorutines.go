package main

import (
	"fmt"
	"runtime"
	"time"
)

func main () {
	defer fmt.Println(1) // потом это 
	defer fmt.Println("1") // сначало это

	// runtime.GOMAXPROCS(3) // устанавлмваем лимит на количество ядер которые может использовать горутина
	
	fmt.Println(runtime.NumCPU()) // горутин может одновременно выполнятся столько же сколько ядер в процессоре, а этой командой можно посмотреть количество ядер
	go printMessage("helo from gorutine") // создание горутины
	time.Sleep(time.Second)
}

func printMessage(msg string) {
	fmt.Println(msg)
}

func ppp (m string) {
	
}






