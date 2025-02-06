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

func CreateList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var list models.List
	json.NewDecoder(r.Body).Decode(&list)
	list.ID = len(lists) + 1
	lists = append(lists, list)
	json.NewEncoder(w).Encode(list)
}

func UpdateList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid user", http.StatusBadRequest)
		return
	}
	for index, item := range lists {
		if item.ID == id {
			lists = append(lists[:index], lists[index+1:]...)
			var list models.List
			item.ID = id
			_ = json.NewDecoder(r.Body).Decode(&list)
			lists = append(lists, list)
			json.NewEncoder(w).Encode(list)
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}

func DeleteList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid user", http.StatusBadRequest)
		return
	}
	for index, item := range lists {
		if item.ID == id {
			lists = append(lists[:index], lists[index:+1]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}
