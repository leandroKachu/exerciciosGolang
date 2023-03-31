package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go CallApi(&wg)
	go CallFunction(&wg)
	go CallJob(&wg)

	wg.Wait()
}

func CallFunction(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("calling function")
	wg.Done()
}

func CallApi(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("call calling API")
	wg.Done()

}

func CallJob(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("call calling job")
	wg.Done()

}
