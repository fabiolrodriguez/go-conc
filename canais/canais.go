package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá Canais", canal)

	fmt.Println("Depois da função escrever")

	for {
		mensagem, aberto := <-canal

		if !aberto {
			break
		}

		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
