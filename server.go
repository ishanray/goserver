package main

import (
  "net/http"
  "log"
)

func main() {
  http.Handle("/", http.FileServer(http.Dir("./public/")))
  log.Println("listening on 8080")
  http.ListenAndServe(":8080", nil)
}
