package handlers

import (
	"API_GO_POSTGRESQL/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, r *http.Request) {
	// Atoi converte string pra inteiro
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		// cria um erro do tipo http com mensagem de erro de servidor e status 500
		// com a ajuda do pacote net/http
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var todo models.Todo // entidade dos models

	err = json.NewDecoder(r.Body).Decode(&todo) // decode do payload json
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		// cria um erro do tipo http com mensagem de erro de servidor e status 500
		// com a ajuda do pacote net/http
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		// cria um erro do tipo http com mensagem de erro de servidor e status 500
		// com a ajuda do pacote net/http
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Erro: %d registros modificados: ", rows)
	}

	resp := map[string]any{

		"Error":   false, // uma forma melhor seria apenas passar o status code com sucesso
		"Message": fmt.Sprintf("Atualizado com sucesso! ID: %d", id),
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
