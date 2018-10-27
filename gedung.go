package main

type Gedung struct {
	Id   int    `form:"id" json:"id"`
	Nama string `form:"nama" json:"nama"`
}

type ResponseGedung struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []Gedung `json:"data"`
}
