package handlers

import (
	"backend/database"
	"backend/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := "SELECT id, name_list, status FROM lists"
	rows, err := database.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var lists []models.List
	for rows.Next() {
		var list models.List
		if err := rows.Scan(&list.ID, &list.Name_list, &list.Status); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		lists = append(lists, list)
	}

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

	var list models.List
	query := "SELECT id, name_list, status FROM lists WHERE id = ?"
	err = database.DB.QueryRow(query, id).Scan(&list.ID, &list.Name_list, &list.Status)
	if err == sql.ErrNoRows {
		http.Error(w, "list not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(list)
}

func CreateList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var list models.List

	if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO lists (name_list, status) VALUES (?, ?)"
	result, err := database.DB.Exec(query, list.Name_list, list.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	list.ID = int(id)

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

	var updatedList models.List
	if err := json.NewDecoder(r.Body).Decode(&updatedList); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	query := "UPDATE lists SET name_list = ?, status = ? WHERE id = ?"
	_, err = database.DB.Exec(query, updatedList.Name_list, updatedList.Status, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updatedList.ID = id
	json.NewEncoder(w).Encode(updatedList)
}

func DeleteList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid list", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM lists WHERE id = ?"
	result, err := database.DB.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "list not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateListStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid list", http.StatusBadRequest)
		return
	}

	query := "UPDATE lists SET status = NOT status WHERE id = ?"
	result, err := database.DB.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "list not found", http.StatusNotFound)
		return
	}

	// Ambil data terbaru setelah update
	var updatedList models.List
	selectQuery := "SELECT id, name_list, status FROM lists WHERE id = ?"
	err = database.DB.QueryRow(selectQuery, id).Scan(&updatedList.ID, &updatedList.Name_list, &updatedList.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedList)
}
