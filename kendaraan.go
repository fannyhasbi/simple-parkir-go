package main

type Kendaraan struct {
	Id         int    `form:"id" json:"id"`
	Plat_nomor string `form:"plat_nomor" json:"plat_nomor"`
}

type ResponseKendaraan struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    []Kendaraan `json:"data"`
}
