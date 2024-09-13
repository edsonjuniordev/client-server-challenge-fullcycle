package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/edsonjuniordev/client-server-challenge-fullcycle/server/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", handlers.CotacaoHandler)

	port := "8080"
	log.Printf("Running server at port: %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
