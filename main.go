package main

import (
  "fmt"
  "net/http"
  "log"
  "strings"
  "io"
)

func main() {
  // NOTE: This should be 80 for production use
  PORT := "8080"

  fmt.Println("Starting server on port " + PORT)

  http.HandleFunc("/", handleRoot)
  log.Fatal(http.ListenAndServe(":" + PORT, nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()

  if strings.Contains(r.URL.Path, "..") {
    io.WriteString(w, "404! Page not found.")
  }

  url_path := r.URL.Path
  if r.URL.Path == "/" {
    url_path = "/index.html"
  }

  

}
