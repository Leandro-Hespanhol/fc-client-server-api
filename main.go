package main

import (
	"fmt"
	"log"
	"modulo-1-go/server"
	"net/http"
)

func main() {
	fmt.Println("Starting server")
	handlers()
	log.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handlers() {
	http.HandleFunc("/cotacao", server.ExchangeHandler)
}
