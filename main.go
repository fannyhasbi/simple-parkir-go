package main

import (
  "fmt"
  "net/http"
)

func handleWriter(w http.ResponseWriter, r *http.Request){
  var message = "Welcome to Parking System"
  w.Write([]byte(message))
}

func main(){
  http.HandleFunc("/", handleWriter)

  var address = "localhost:4040"
  fmt.Printf("Server running on %s", address)
  err := http.ListenAndServe(address, nil)

  if err != nil {
    fmt.Println(err.Error())
  }
}