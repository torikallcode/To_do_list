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
	selectQuery := "SELECT id, name_list, status FROM list WHERE id = ?"
	err = database.DB.QueryRow(selectQuery, id).Scan(&updateList.ID, &updateList.Name_list, &updateList.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updateList)
}
