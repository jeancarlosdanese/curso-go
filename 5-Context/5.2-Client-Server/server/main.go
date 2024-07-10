package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)

}

func handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(5 * time.Second):
		// Imprime no console que a requisição foi processada
		log.Println("Request processada com sucesso!")
		// Imprime no Browser que a requisição foi processada
		w.Write([]byte("Request processada com sucesso!"))
	case <-ctx.Done():
		// Imprime no console que a requisição foi cancelada
		log.Println("Request cancelada pelo cliente!")
	}
}
