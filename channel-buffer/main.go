package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int, 100)
	// nesse exemplo, criei um cara chamado chan = channel e passei a quantidade de instance dele = 100
	go setList(channel)

	for v := range channel {
		fmt.Println("Enviando:", v)
		time.Sleep(time.Second)
	}
}

func setList(channel chan<- int) {
	for i := 0; i <= 100; i++ {
		channel <- i
		fmt.Println("Recebendo:", i)
	}

	close(channel)
}
