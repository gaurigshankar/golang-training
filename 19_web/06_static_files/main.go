package main

import "net/http"

// curl -s http://localhost:9999/static/css/styles.css

func main() {
  fs := http.FileServer(http.Dir("assets/"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))
  http.ListenAndServe(":9999", nil)
}
