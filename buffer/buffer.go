package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Abra seu canal para o mundo!"
	canal <- "Go horse"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}
