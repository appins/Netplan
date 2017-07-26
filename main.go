package main

import (
  "fmt"
  "net"
  "net/http"
  "log"
  "strings"
  "io"
  "os"
  "strconv"
)

// Create a map for counting the amount of requests
var reqCount map[string]int

// Function for checking if file or folder exists
func pathExists(path string) bool {
  _, err := os.Stat(path)
  if err == nil {
    return true
  }
  if os.IsNotExist(err) {
    return false
  }
  return true
}

func main() {
  reqCount = make(map[string]int)

  // NOTE: This should be changed to your ip when testing so you don't reach the limit
  reqCount["10.0.0.188"] = -1

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

  // TODO: Send a 404 header and create a custom page in public
  dat, err := os.Open("public" + url_path)
  if err != nil {
    ip, _, _ := net.SplitHostPort(r.RemoteAddr)
    fmt.Println("The user (" + ip + ") made a request for " + url_path + ", but there was nothing there!")
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

  ip, _, _ := net.SplitHostPort(r.RemoteAddr)
  if reqCount[ip] > 100 {
    io.WriteString(w, "var userid = \"This ip has created too many id's!\"")
    fmt.Println("A user (" + ip + ") has created over 100 id's")
    return
  }

  // Notice that if the int is under zero, than we don't increment it
  if reqCount[ip] >= 0 {
    reqCount[ip]++
  }

  // Create a random username and check if it exists
  userid := getRandom()
  for i := 0; pathExists("./entries/" + userid); i++ {
    if i > 100 {
      fmt.Println("After 100 tries, we couldn't fint an ID for " + ip)
      fmt.Println("Please reset the entires folder or make the random function better")
      // IDEA: Send the user a script that will alert them about the issue
      return
    }
    userid = getRandom()
  }

  os.MkdirAll("./entries/" + userid, 0777)
  jsfile := "var userid = \"" + userid + "\";"

  w.Header().Add("Content-Type", "applictaion/javascript")
  io.WriteString(w, jsfile)
}

func handleJournal(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()

  if strings.Contains(r.URL.Path, "..") {
    io.WriteString(w, "404! Page not found.")
  }

  path := strings.Split(r.URL.Path, "/")[2]
  if !pathExists("./entries/" + path) {
    dat, err := os.Open("./public/notfound.html")
    if err != nil {
      fmt.Println("public/notfound.html is missing!")
      io.WriteString(w, "The journal you are looking for was not found.")
      return
    }
    w.Header().Add("Content-Type", "text/html")
    io.Copy(w, dat)
    return
  }

  journal_url := strings.Split(r.URL.Path, "/")[3]

  if journal_url == "" {
    journal_url = "index.html"
  }

  // The handler when an entry is requested or written too
  if journal_url == "entry" || journal_url == "entryedit" {
    entNum := strings.Split(r.URL.Path + "/", "/")[4]
    entryNum, err := strconv.Atoi(entNum)
    if err != nil || entryNum > 10000 {
      io.WriteString(w, "Journal entry number is invalid.")
      return
    }

    entryExists := pathExists("./entries/" + path + "/" + entNum + ".ent")

    if entryExists && journal_url == "entry" {
      dat, err := os.Open("./entries/" + path + "/" + entNum + ".ent")
      if err != nil {
        io.WriteString(w, "ERROR!")
        fmt.Println("Couldn't open a path even though it was supposed to exist")
        fmt.Println("( path=" + path + ", entNum=" + entNum + ")")
        return
      }
      io.Copy(w, dat)
      return
    }
    if !entryExists {

    }

  }

  dat, err := os.Open("./journal_static/" + journal_url)
  if err != nil {
    io.WriteString(w, "404! Page not found.")
  }

  var contentType string
  fileExt := strings.Split(journal_url + ".", ".")[1]

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
