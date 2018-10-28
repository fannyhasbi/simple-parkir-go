package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/gedung", returnGedung).Methods("GET")
	router.HandleFunc("/kendaraan", returnKendaraan).Methods("GET")

	// petugas
	router.HandleFunc("/petugas", returnPetugas).Methods("GET")
	router.HandleFunc("/petugas", returnLoginPetugas).Methods("POST")

	http.Handle("/", router)

	fmt.Println("Running on port 4040")
	log.Fatal(http.ListenAndServe(":4040", router))
}
