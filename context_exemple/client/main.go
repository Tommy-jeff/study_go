package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// for i := 0; i < 1000; i++ {
		// Criação de um contexto com timeout de 10 segundos para cada requisição
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		// Criação da requisição HTTP com o contexto
		req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/", nil)
		if err != nil {
			log.Fatal(err)
		}
	
		// log.Printf("Enviando requisição %d", i)
		log.Printf("Enviando requisição")

		// Envio da requisição HTTP
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("Erro na requisição: %v", err)
			return
		}
		// Leitura e exibição da resposta
		defer res.Body.Close()
		io.Copy(os.Stdout, res.Body)
	// }
}