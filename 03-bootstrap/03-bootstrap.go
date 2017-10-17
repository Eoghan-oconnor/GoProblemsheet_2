//Eoghan O'Connor
//Go Problem Sheet 
//03-bootstrap
//Adapted From: https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server

package main

import (
  "log"
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("03-bootstrap"))
  http.Handle("/", fs)

  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}