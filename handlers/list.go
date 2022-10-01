package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hugovallada/go-api-postgres/models"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao buscar os registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
