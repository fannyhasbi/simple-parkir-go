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

	for rows.Next() {
		if err := rows.Scan(&gedung.Id, &gedung.Nama); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_gedung = append(arr_gedung, gedung)
		}
	}

	response.Status = 200
	response.Message = "OK"
	response.Data = arr_gedung

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func returnKendaraan(w http.ResponseWriter, h *http.Request) {
	var kendaraan Kendaraan
	var arr_kendaraan []Kendaraan
	var response ResponseKendaraan

	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM kendaraan")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&kendaraan.Id, &kendaraan.Plat_nomor); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_kendaraan = append(arr_kendaraan, kendaraan)
		}
	}

	response.Status = 200
	response.Message = "OK"
	response.Data = arr_kendaraan

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func returnPetugas(w http.ResponseWriter, r *http.Request) {
	var petugas Petugas
	var arr_petugas []Petugas
	var response ResponsePetugas

	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, nama FROM petugas")
	if err != nil {
		log.Println("2")
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&petugas.Id, &petugas.Nama); err != nil {
			log.Println("2")
			log.Fatal(err.Error())
		} else {
			arr_petugas = append(arr_petugas, petugas)
		}
	}

	response.Status = 200
	response.Message = "OK"
	response.Data = arr_petugas

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
