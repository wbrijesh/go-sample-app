package main

import (
  "fmt"
  "net/http"
  "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World!")
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./hello.html")
}

func main() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/html", htmlHandler)
  
  http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
