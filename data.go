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
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&petugas.Id, &petugas.Nama); err != nil {
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

func returnLoginPetugas(w http.ResponseWriter, r *http.Request) {
	var petugas LoginPetugas
	var response ResponseLoginPetugas

	username := r.FormValue("username")
	password := r.FormValue("password")

	db := connect()
	defer db.Close()

	err := db.QueryRow("SELECT id, username, password, nama FROM petugas WHERE username = ?", username).Scan(&petugas.Id, &petugas.Username, &petugas.Password, &petugas.Nama)

	if petugas.Password == password {
		response.Status = 401
		response.Message = "Unauthorized"
		response.Data = ""

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		if err != nil {
			response.Status = 401
			response.Message = "Unauthorized"
			response.Data = ""
		} else {
			response.Status = 200
			response.Message = "OK"
			response.Data = petugas.Nama
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func returnKendaraanMasuk(w http.ResponseWriter, r *http.Request) {
	var response ResponseKendaraanMasuk

	id_petugas := r.FormValue("id_petugas")
	id_kendaraan := r.FormValue("id_kendaraan")

	db := connect()
	defer db.Close()

	_, err := db.Exec("INSERT INTO kendaraan_masuk (id_petugas, id_kendaraan) VALUES (?, ?)", id_petugas, id_kendaraan)
	if err != nil {
		http.Error(w, "Server error, unable to insert data", 500)
		return
	}

	response.Status = 200
	response.Message = "OK"
	response.Data = "Insert successful"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
