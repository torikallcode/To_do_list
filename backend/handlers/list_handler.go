package handlers

import (
	"backend/database"
	"backend/models"
	"encoding/json"
	"net/http"
)

var lists []models.List

func GetAllList(w http.ResponseWriter, r *http.Request) {
	// set header response sebaga json
	w.Header().Set("Content-Type", "application/json")

	// membuat kode query yang akan dieksekusi
	query := "SELECT id, name_list, status FROM list"

	// digunakan untuk mengeksekusi query ke database. (database.DB adalah tempat koneksi databasenya)
	rows, err := database.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//menutup cursor yang digunakan untuk mengakses hasil query (setelah fungsi selesai)
	defer rows.Close()

	// membaca baris hasil query satu per satu
	for rows.Next() {
		var list models.List
		// memindahkan data dari baris saat ini(di database) ke variable list yang sudah dibuat
		if err := rows.Scan(&list.ID, &list.Name_list, &list.Status); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// menambahkan baris saat ini ke slice lists
		lists = append(lists, list)
	}
	// ubah format slice licsts ke json dan kirim sebagai response
	json.NewEncoder(w).Encode(lists)
}
