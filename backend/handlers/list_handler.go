package handlers

import (
	"backend/models"
	"encoding/json"
	"net/http"
)

var lists []models.List

func GetAllList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lists)
}
