package goroutines

import (
	"fmt"
	"time"
)

// Função que processa trabalhos ela recebe um id de trabalhador, um canal de trabalhos e um canal de resultados
// Ela lê trabalhos do canal de trabalhos, processa-os (neste caso, apenas multiplica por 2) e envia os resultados para o canal de resultados
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
		time.Sleep(time.Second * 2) // Simula um trabalho demorado
        fmt.Printf("Worker %d processando %d\n", id, j)
        results <- j * 2
    }
}

func SuperExemple() {

	// Criando canais para trabalhos e resultados
    jobs := make(chan int, 5)
    results := make(chan int, 5)

	// Iniciando 3 goroutines de trabalhadores, ou seja, 3 workers simultâneos
    for w := 1; w <= 5; w++ {
        go worker(w, jobs, results)
    }

	// Enviando trabalhos para os trabalhadores
    for j := 1; j <= 5; j++ {
        jobs <- j
    }

	// Fechando o canal de trabalhos
    close(jobs)

	// Coletando os resultados
    for r := 1; r <= 5; r++ {
        fmt.Println("Resultado:", <-results)
    }
}
