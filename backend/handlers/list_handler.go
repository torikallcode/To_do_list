package handlers

import (
	"backend/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var lists []models.List

func GetAllList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lists)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid user", http.StatusBadRequest)
		return
	}
	for _, item := range lists {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}
