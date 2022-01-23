package main

import (
	"fmt"
	"time"
)

func main() {
	// Concorrencia != Paralelismo
	go escrever("Oía zé") //goroutine
	escrever("teste concorrencia")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
