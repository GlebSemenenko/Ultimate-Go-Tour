package main

import (
	"fmt"
	"os/exec"
	"time"
	"runtime"
)

// программа выключает пк через время записанное в t
func main() {
    var t int64 = 3 // выбираем время
    fmt.Println(t, " sec to sleep")
    dura := time.Duration(t) * time.Second // В time есть  type Duration поэтому время 
    time.Sleep(dura) // устанавливаем таймер 

    if runtime.GOOS == "windows" { // определяем ос
        cmd := exec.Command("shutdown", "/s", "/t", "0")
        err := cmd.Run() // запуск консольной команды
        if err != nil {
            fmt.Println("Ошибка при выполнении команды:", err)
        } else {
            fmt.Println("Команда выполнена успешно. ПК выключится.")
        }
    } else {
        fmt.Println("Эта программа предназначена для Windows.")
    }
}
