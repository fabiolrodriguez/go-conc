package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup

	waitgroup.Add(4)

	go func() {
		escrever("Oía zé") //waitgroup
		waitgroup.Done()   // -1
	}()

	go func() {
		escrever("teste concorrencia")
		waitgroup.Done() // -1
	}()

	go func() {
		escrever("goroutine 3")
		waitgroup.Done() // -1
	}()

	go func() {
		escrever("goroutine 4")
		waitgroup.Done() // -1
	}()

	waitgroup.Wait() // 1 - para quando chegar em 0

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
