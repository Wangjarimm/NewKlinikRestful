package main

import (
	"belajar/controller/auth"
	"belajar/controller/dokter"
	"belajar/controller/pasien"
	"belajar/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	database.InitDB()
	fmt.Println("Hello World")

	router := mux.NewRouter()

	router.HandleFunc("/regis", auth.Registration).Methods("POST")
	router.HandleFunc("/login", auth.Login).Methods("POST")

	// Router handler pasien
	router.HandleFunc("/pasien", pasien.GetPasien).Methods("GET")
	router.HandleFunc("/pasien", auth.JWTAuth(pasien.PostPasien)).Methods("POST")
	router.HandleFunc("/pasien/{id}", auth.JWTAuth(pasien.PutPasien)).Methods("PUT")
	router.HandleFunc("/pasien/{id}", auth.JWTAuth(pasien.DeletePasien)).Methods("DELETE")

	// Router handler dokter
	router.HandleFunc("/dokter", dokter.GetDokter).Methods("GET")
	router.HandleFunc("/dokter", auth.JWTAuth(dokter.PostDokter)).Methods("POST")
	router.HandleFunc("/dokter/{id}", auth.JWTAuth(dokter.PutDokter)).Methods("PUT")
	router.HandleFunc("/dokter/{id}", auth.JWTAuth(dokter.DeleteDokter)).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		Debug:          true,
	})

	handler := c.Handler(router)

	fmt.Println("Server is running on http://localhost:4500")
	log.Fatal(http.ListenAndServe(":4500", handler))

}
