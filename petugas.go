package main

type Petugas struct {
	Id   int    `form:"id" json:"id"`
	Nama string `form:"nama" json:"nama"`
}

type LoginPetugas struct {
	Id       int    `form:"id" json:"id"`
	Nama     string `form:"nama" json:"nama"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type ResponsePetugas struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Petugas `json:"data"`
}

type ResponseLoginPetugas struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
