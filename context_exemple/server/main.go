package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Print("iniciou a chamada home")
	defer log.Print("Finalizou a requisição")

	select {
	case <-time.After(5 * time.Second):
		log.Print("Requisição processada com sucesso")
		w.Write([]byte("Hello, World!\n"))
	case <-ctx.Done():
		log.Print("Requisição cancelada pelo cliente")
		http.Error(w, "Requisição cancelada pelo cliente", http.StatusRequestTimeout)
	}
}
