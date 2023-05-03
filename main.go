package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/", helloWorld)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, world!")
}