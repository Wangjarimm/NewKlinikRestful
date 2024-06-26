package pasien

import (
    "encoding/json"
    "net/http"
    "strconv"

    "belajar/database"
    "github.com/gorilla/mux"
    "belajar/model/pasien"
)

func GetPasien(w http.ResponseWriter, r *http.Request) {
    rows, err := database.DB.Query("SELECT * FROM pasien")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var pasiens []pasien.Pasien
    for rows.Next() {
        var c pasien.Pasien
        if err := rows.Scan(&c.Id, &c.Namalengkap, &c.Nik, &c.Jeniskelamin, &c.Tempatlahir, &c.Tanggallahir, &c.Alamat, &c.Nohp); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        pasiens = append(pasiens, c)
    }

    if err := rows.Err(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(pasiens); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}


func PostPasien(w http.ResponseWriter, r *http.Request) {
	var pc pasien.Pasien
	if err := json.NewDecoder(r.Body).Decode(&pc); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Prepare the SQL statement for inserting a new pasien
	query := `
		INSERT INTO pasien (nama_lengkap, nik, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat, no_hp) 
		VALUES (?, ?, ?, ?, ?, ?, ?,)`

	// Execute the SQL statement
	res, err := database.DB.Exec(query, pc.Namalengkap, pc.Nik, pc.Jeniskelamin, pc.Tempatlahir, pc.Tanggallahir, pc.Alamat, pc.Nohp)
	if err != nil {
		http.Error(w, "Failed to insert course: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the last inserted ID
	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve last insert ID: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the newly created ID in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Pasien added successfully",
		"id":      id,
	})
}

func PutPasien(w http.ResponseWriter, r *http.Request) {
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
	var pc pasien.Pasien
	if err := json.NewDecoder(r.Body).Decode(&pc); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Prepare the SQL statement for updating the category admin
	query := `
		UPDATE pasien 
		SET nama_lengkap=?, nik=?, jenis_kelamin=?, tempat_lahir=?, tanggal_lahir=?, alamat=?, no_hp=?
		WHERE id = ?`
	// Execute the SQL statement
	result, err := database.DB.Exec(query, pc.Namalengkap, pc.Nik, pc.Jeniskelamin, pc.Tempatlahir, pc.Tanggallahir, pc.Alamat, pc.Nohp, id)
	if err != nil {
		http.Error(w, "Failed to update course: "+err.Error(), http.StatusInternalServerError)
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
		"message": "Pasien updated successfully",
	})
}

func DeletePasien(w http.ResponseWriter, r *http.Request) {
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
		DELETE FROM pasien
		WHERE id = ?`

	// Execute the SQL statement
	result, err := database.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Failed to delete course: "+err.Error(), http.StatusInternalServerError)
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
		"message": "Pasien deleted successfully",
	})
}

