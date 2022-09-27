package main

import (
	"fmt"
	"time"
)

func procesar(i int, c chan int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Termina")
	c <- i
}

func main() {
	canal := make(chan int) //similar a signal C
	for i := 0; i < 10; i++ {
		go procesar(i, canal)
	}
	var lectura int
	for i := 0; i < 10; i++ {
		lectura = <-canal
		//<-canal //también es valido solo llamarlo, igual espera
	}
	fmt.Println("Terminamos", lectura)
}
