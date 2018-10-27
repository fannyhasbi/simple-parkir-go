package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func returnGedung(w http.ResponseWriter, r *http.Request) {
	var gedung Gedung
	var arr_gedung []Gedung
	var response ResponseGedung

	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM gedung")
	if err != nil {
		log.Print(err)
	}

	log.Println(rows)

	for rows.Next() {
		if err := rows.Scan(&gedung.Id, &gedung.Nama); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_gedung = append(arr_gedung, gedung)
		}
	}

	response.Status = 1
	response.Message = "OK"
	response.Data = arr_gedung

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
