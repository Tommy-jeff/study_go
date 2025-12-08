package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	// Configuração do servidor HTTP
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

// Manipulador para a rota raiz
func home(w http.ResponseWriter, r *http.Request) {
	// Obtém o contexto da requisição
	ctx := r.Context()
	log.Print("iniciou a chamada home")
	defer log.Print("Finalizou a requisição")

	select {
	// Simula um processamento que leva 1 segundo
	case <-time.After(time.Second):
		log.Print("Requisição processada com sucesso")
		w.Write([]byte("Hello, World!\n"))
	// Se o contexto for cancelado (por exemplo, se o cliente desconectar), essa case será executada
	case <-ctx.Done():
		log.Print("Requisição cancelada pelo cliente")
		http.Error(w, "Requisição cancelada pelo cliente", http.StatusRequestTimeout)
	}
}
