package handlers

import (
	"backend/models"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

// Gunakan mutex untuk thread-safety
var (
	lists = []models.List{
		{ID: 1, Name_list: "Program", Status: true},
	}
	mu     sync.Mutex
	nextID = 2 // Mulai dari 2 karena sudah ada item dengan ID 1
)

func GetAllList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(lists)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid list", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	for _, item := range lists {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "list not found", http.StatusNotFound)
}

func CreateList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var list models.List

	if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Gunakan nextID untuk memastikan ID unik
	list.ID = nextID
	nextID++

	lists = append(lists, list)
	json.NewEncoder(w).Encode(list)
}

func UpdateList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid list", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for index, item := range lists {
		if item.ID == id {
			var updatedList models.List
			if err := json.NewDecoder(r.Body).Decode(&updatedList); err != nil {
				http.Error(w, "invalid input", http.StatusBadRequest)
				return
			}

			// Pertahankan ID asli
			updatedList.ID = id

			// Ganti item di slice
			lists[index] = updatedList

			json.NewEncoder(w).Encode(updatedList)
			return
		}
	}
	http.Error(w, "list not found", http.StatusNotFound)
}

func DeleteList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid list", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for index, item := range lists {
		if item.ID == id {
			// Gunakan metode slice yang benar untuk menghapus
			lists = append(lists[:index], lists[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "list not found", http.StatusNotFound)
}

func UpdateListStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid list", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for index, item := range lists {
		if item.ID == id {
			lists[index].Status = !item.Status
			json.NewEncoder(w).Encode(lists[index])
			return
		}
	}
	http.Error(w, "list not found", http.StatusNotFound)
}
