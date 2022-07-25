package handlers

import (
	"API_GO_POSTGRESQL/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Get(w http.ResponseWriter, r *http.Request) {

	// ===== PARSE DO ID =====
	// Atoi converte string pra inteiro
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		// cria um erro do tipo http com mensagem de erro de servidor e status 500
		// com a ajuda do pacote net/http
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// ===== MODEL GET =====
	todo, err := models.Get(int64(id))
	if err != nil {
		log.Printf("Erro ao obter registro %d: %v", id, err)
	}

	// ===== HTTP RESPONSE
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
