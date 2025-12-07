package goroutines

import (
	"fmt"
	"time"
)

//Thread 1
func ChannelsExemple() {

	//criando um canal
	canal := make(chan string)

	//Thread 2
	go func() {
		//enviando mensagem para o canal
		canal <- "Olá do canal!"
	}()

	mensagem := <-canal

	println(mensagem)
}


func ChannelsExemple2() {

	canal := make(chan string)

	go func() {
		canal <- "Olá do canal!"
	}()

	// This select statement waits on multiple communication operations.
	select {
		case mensagem := <-canal:
			println(mensagem)
		case <-time.After(2 * time.Second):
			println("Tempo esgotado esperando a mensagem do canal.")
	}

}
func ChannelsExemple3() {

	queue := make(chan int)

	go func () {
		for i := 0; i < 99999; i++ {
			queue <- i
		}
		close(queue)
	}()

	for x := range queue {
		fmt.Println(x)
	}
}