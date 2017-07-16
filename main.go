package main

import (
  "fmt"
  "net/http"
  "log"
  "strings"
  "io"
  "os"
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

  dat, err := os.Open("public" + url_path)
  if err != nil {
    fmt.Println("The user made a request for " + url_path + ", but there was nothing there")
    io.WriteString(w, "404! Page not found.")
    return
  }

  io.Copy(w, dat)

  var contextType string
  fileExt := strings.Split(url_path, ".")[1]

  switch fileExt {
  case "css":
    contextType = "text/css"
  case "html":
    contextType = "text/html"
  case "js":
    contextType = "application/javascript"
  default:
    contextType = "text/plain"
  }

  w.Header().Add("Context-Type", contextType)
  io.Copy(w, dat)

}
