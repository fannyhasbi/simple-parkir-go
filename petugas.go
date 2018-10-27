package main

type Petugas struct {
	Id   int    `form:"id" json:"id"`
	Nama string `form:"nama" json:"nama"`
}

type ResponsePetugas struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Petugas `json:"data"`
}
