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
	// Set header response sebagai JSON
	// Eksekusi query SELECT untuk mengambil semua data
	// Membaca hasil query baris per baris
	// Menyimpan setiap baris ke dalam slice lists
	// Mengirim seluruh data lists sebagai JSON response

	w.Header().Set("Content-Type", "application/json")

	var lists []models.List

	query := "SELECT id, name_list, status FROM lists"
	rows, err := database.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

	// 	Set header response sebagai JSON
	// Ambil parameter ID dari URL
	// Konversi ID dari string ke integer
	// Eksekusi query SELECT untuk mengambil data spesifik
	// Memindahkan data ke variabel list
	// Mengirim data list sebagai JSON response

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid list", http.StatusBadRequest)
		return
	}

	var list models.List
	query := "SELECT id, name_list, status FROM lists WHERE id = ?"
	err = database.DB.QueryRow(query, id).Scan(&list.ID, &list.Name_list, &list.Status)
	if err == sql.ErrNoRows {
		http.Error(w, "rows not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(list)
}

func CreateList(w http.ResponseWriter, r *http.Request) {

	// Set header response sebagai JSON
	// Decode data JSON dari request body
	// Validasi input
	// Eksekusi query INSERT ke database
	// Ambil ID dari baris yang baru dimasukkan
	// Kirim data list yang baru dibuat sebagai response

	w.Header().Set("Content-Type", "application/json")

	var list models.List
	if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
		http.Error(w, "Invalid list", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO lists (name_list, status) VALUE (?, ?)"
	result, err := database.DB.Exec(query, &list.Name_list, &list.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	list.ID = int(id)

	json.NewEncoder(w).Encode(list)

}

func UpdateList(w http.ResponseWriter, r *http.Request) {
	// Set header response sebagai JSON
	// Ambil parameter ID dari URL
	// Konversi ID dari string ke integer
	// Decode data JSON dari request body
	// Validasi input
	// Eksekusi query UPDATE ke database
	// Kirim data list yang diupdate sebagai response

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid list", http.StatusBadRequest)
		return
	}

	var updateList models.List
	if err := json.NewDecoder(r.Body).Decode(&updateList); err != nil {
		http.Error(w, "invalid list", http.StatusBadRequest)
		return
	}

	query := "UPDATE lists SET name_list = ?, status = ? WHERE id = ? "
	_, err = database.DB.Exec(query, &updateList.Name_list, &updateList.Status, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updateList.ID = id

	json.NewEncoder(w).Encode(updateList)
}

func DeleteList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //set header sebagai json
	params := mux.Vars(r)                              // ambil parameter(id) dari url
	id, err := strconv.Atoi(params["id"])              // konversi id ke int
	if err != nil {
		http.Error(w, "invalid list", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM lists WHERE id = ?"  // kode query untuk menghapus data dari tabel list berdasarkan id
	result, err := database.DB.Exec(query, id) // mengeksekusi query sql dengan parameter id
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected() //(result.RowsAffected()) method yang digunakan untuk mengembalika jumlah baris yang terpengaruh oleh query
	if rowsAffected == 0 {
		http.Error(w, "list not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // jika query berhasil dan data berhasil dihapus maka akan mengirim response (204 no content)
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

	var updateList models.List
	selectQuery := "SELECT id, name_list, status FROM lists WHERE id = ?"
	err = database.DB.QueryRow(selectQuery, id).Scan(&updateList.ID, &updateList.Name_list, &updateList.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updateList)
}
