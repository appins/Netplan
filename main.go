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
  http.HandleFunc("/makenew.js", handleNew)
  http.HandleFunc("/journal/", handleJournal)
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
    fmt.Println("The user made a request for " + url_path + ", but there was nothing there!")
    io.WriteString(w, "404! Page not found.")
    return
  }

  var contentType string
  fileExt := strings.Split(url_path, ".")[1]

  switch fileExt {
  case "css":
    contentType = "text/css"
  case "html":
    contentType = "text/html"
  case "js":
    contentType = "application/javascript"
  default:
    contentType = "text/plain"
  }

  w.Header().Add("Content-Type", contentType)
  io.Copy(w, dat)
}

func handleNew(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()

  userid := getRandom()
  os.MkdirAll("./entries/" + userid, 0777)
  jsfile := "var userid = \"" + userid + "\";"

  io.WriteString(w, jsfile)

  w.Header().Add("Content-Type", "applictaion/javascript")
}

func handleJournal(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()


}
