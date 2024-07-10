package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		// Some code that may panic
		panic("Something went wrong!")
	})

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", recoverMiddleware(mux)); err != nil {
		fmt.Println(err)
	}
}

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("recovered panic: %v", r)
				// debug.PrintStack()
				http.Error(w, "Something went wrong!", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
