package handlers

import (
	"API_GO_POSTGRESQL/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo // entidade dos models

	err := json.NewDecoder(r.Body).Decode(&todo) // decode do payload json
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		// cria um erro do tipo http com mensagem de erro de servidor e status 500
		// com a ajuda do pacote net/http
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true, // uma forma melhor seria apenas passar o status code com erro
			"Message": fmt.Sprintf("Erro ao inserir: %v", err),
		}
	} else {
		resp = map[string]any{

			"Error":   false, // uma forma melhor seria apenas passar o status code com sucesso
			"Message": fmt.Sprintf("Inserido com sucesso! ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
