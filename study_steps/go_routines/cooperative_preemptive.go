package goroutines

import (
	"fmt"
	"runtime"
	"time"
)

// Processamento cooperativo significa que as goroutines devem ceder o controle voluntariamente para permitir que outras goroutines sejam executadas.
// Diferente do preemptivo, onde o sistema operacional pode interromper uma thread a qualquer momento para dar lugar a outra.
// Neste exemplo vemos como o Go pode agir de forma preemtiva, mesmo sendo cooperativo por natureza.

func CooperativePreemptive() {

	runtime.GOMAXPROCS(1)
	fmt.Println("Iniciando processamento cooperativo com preempção...")

	go func() {
		for {

		}
	}()
	
	time.Sleep(time.Second)
	fmt.Println("Esta mensagem será exibida mesmo com a goroutine em loop infinito.")
}