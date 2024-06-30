package dokter

import (
	"encoding/json"
	"net/http"
	"strconv"

	"belajar/database"
	"belajar/model/dokter"

	"github.com/gorilla/mux"
)

func GetDokter(w http.ResponseWriter, r *http.Request) {
    rows, err := database.DB.Query("SELECT * FROM dokter")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var dokters []dokter.Dokter
    for rows.Next() {
        var s dokter.Dokter
        if err := rows.Scan(&s.Id,&s.Nid,&s.Nama,&s.Keahlian,&s.Nohp); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        dokters = append(dokters, s)
    }

    if err := rows.Err(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(dokters)
}

func PostDokter(w http.ResponseWriter, r *http.Request) {
	var ps dokter.Dokter
	if err := json.NewDecoder(r.Body).Decode(&ps); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Query untuk memasukan mahasiswa ke dalam table
	query := `
		INSERT INTO dokter (nid, nama, keahlian, no_hp) 
		VALUES (?, ?,?,?)`

	// Mengeksekusi query
	res, err := database.DB.Exec(query, ps.Nid, ps.Nama, ps.Keahlian, ps.Nohp)
	if err != nil {
		http.Error(w, "Failed to insert Dokter: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Ambil id terakhir
	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve last insert ID: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the newly created ID in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Dokter added successfully",
		"id":      id,
	})
}

func PutDokter(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari URL
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Decode JSON body
	var ps dokter.Dokter
	if err := json.NewDecoder(r.Body).Decode(&ps); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Query ubah mahasiswa
	query := `
		UPDATE dokter 
		SET nid=?, nama=?, keahlian=?, no_hp=?
		WHERE id=?`

	// Execute the SQL statement
	result, err := database.DB.Exec(query,ps.Nid, ps.Nama, ps.Keahlian, ps.Nohp, id)
	if err != nil {
		http.Error(w, "Failed to update dokter: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to retrieve affected rows: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if any rows were updated
	if rowsAffected == 0 {
		http.Error(w, "No rows were updated", http.StatusNotFound)
		return
	}

	// Return success message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "dokter updated successfully",
	})
}

func DeleteDokter(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Prepare the SQL statement for deleting a category admin
	query := `
		DELETE FROM dokter
		WHERE id = ?`

	// Execute the SQL statement
	result, err := database.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Failed to delete dokter: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to retrieve affected rows: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "No rows were deleted", http.StatusNotFound)
		return
	}

	// Return the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "dokter deleted successfully",
	})
}

