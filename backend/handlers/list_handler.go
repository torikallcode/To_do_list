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

func GetList(w http.ResponseWriter, r *http.Request) {
	// Set header sebagi json
	w.Header().Set("Content-Type", "application/json")
	// ambil parameter url (id)
	params := mux.Vars(r)
	// konversi id ke int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid list", http.StatusBadRequest)
		return
	}

	// buat variable sebelum menyimpan list
	var list models.List
	// buat kode query yang akan di eksekusi di sql untuk mengambil data dari tabel list
	query := "SELECT id, name_list, status FROM list WHERE id = ?"
	// (QueryRow = mengambil 1 baris hasil), ()
	err = database.DB.QueryRow(query, id).Scan(&list.ID, &list.Name_list, &list.Status)
	if err == sql.ErrNoRows {
		http.Error(w, "Rows not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ubah format list ke json dan kirim sebagai response
	json.NewEncoder(w).Encode(list)
}

func CreateList(w http.ResponseWriter, r *http.Request) {
	// set header dalam format json
	w.Header().Set("Content-Type", "application/json")
	// buat variable untuk menyimpan data yang di decode
	var list models.List
	// decode body
	if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
		http.Error(w, "invalid list", http.StatusNotFound)
		return
	}

	// buat kode query untuk memasukkan data baru ke tabel list
	query := "INSERT INTO lists (name_list, status) VALUE (?, ?)"
	// mengeksekusi query dengan parameter yang digunakan
	result, err := database.DB.Exec(query, list.Name_list, list.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// mengembalika ID dari baris yang baru saja dimasukkan
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// menyimpan id yang dimasukkan ke database tadi ke list.id
	list.ID = int(id)
	// ubah dalam format json dan kirim sebagai response
	json.NewEncoder(w).Encode(list)
}

func UpdateList(w http.ResponseWriter, r *http.Request) {
	// set header sebagai json
	w.Header().Set("Content-Type", "application/json")
	// ambil parameter url (id)
	params := mux.Vars(r)
	// konversi id ke int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid list", http.StatusBadRequest)
		return
	}

	// buat variable untuk menampung data yang di decode
	var UpdateList models.List
	// decode data
	if err := json.NewDecoder(r.Body).Decode(&UpdateList); err != nil {
		http.Error(w, "Invalid list", http.StatusBadRequest)
		return
	}

	// buat kode query untuk update data yang akan diupdate
	query := "UPDATE list SET name_list = ?, status = ?, id = ?"
	//mengeksekusi query dengan parameter yang digunakan
	_, err = database.DB.Exec(query, UpdateList.Name_list, UpdateList.Status, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// memberika id terbaru pada update list
	UpdateList.ID = id
	// mengubah updatelist ke dalam json dann mengirimnya sebagai response
	json.NewEncoder(w).Encode(UpdateList)
}
