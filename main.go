package main

import (
	"API_GO_POSTGRESQL/configs"
	"API_GO_POSTGRESQL/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	// ===== LOAD CONFIGURATIONS =====
	err := configs.Load()
	if err != nil {
		log.Printf("Erro: Não foi possível carregar as configurações: %v", err)
		panic(err) // mata aplicação se não for possível carregar as configurações
	}

	// ===== ROUTER =====
	r := chi.NewRouter()
	r.Post("/", handlers.Create)       // apontar a função
	r.Get("/", handlers.List)          // apontar a função
	r.Get("/{id}", handlers.Get)       // apontar a função
	r.Put("/{id}", handlers.Update)    // apontar a função
	r.Delete("/{id}", handlers.Delete) // apontar a função

	// ===== LISTEN ======
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
