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

// Create a map for storing the last journal entry a user requested
var lastEntryNum map[string]string

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
  lastEntryNum = make(map[string]string)

  // NOTE: This should be changed to your ip when testing so you don't reach the limit
  reqCount["10.0.0.188"] = -1

  // NOTE: This should be 80 for production use
  PORT := "80"

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
    fmt.Println("A user (" + ip + ") made a request for " + url_path + ", but there was nothing there!")
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
  for i := 0; pathExists("./entries/" + userid + ".json"); i++ {
    if i > 100 {
      fmt.Println("After 100 tries, we couldn't fint an ID for " + ip)
      fmt.Println("Please reset the entires folder or make the random function better")
      // IDEA: Send the user a script that will alert them about the issue
      return
    }
    userid = getRandom()
  }

  lastEntryNum[userid] = "1";

  os.MkdirAll("./entries/", 0777)
  newJournal(userid)
  jsfile := "var userid = \"" + userid + "\";"

  w.Header().Add("Content-Type", "applictaion/javascript")
  io.WriteString(w, jsfile)
}

func handleJournal(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()

  if strings.Contains(r.URL.Path, "..") {
    io.WriteString(w, "404! Page not found.")
  }

  path := strings.ToLower(strings.Split(r.URL.Path, "/")[2])
  if !pathExists("./entries/" + path + ".json") {
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
  if journal_url == "entry" || journal_url == "entryedit" || journal_url == "title" || journal_url == "titleedit" {
    // Add another '/' at the end so that we don't get index errors (as often)
    entNum := strings.Split(r.URL.Path + "/", "/")[4]
    entryNum, err := strconv.Atoi(entNum)
    if err != nil || entryNum > 2500 || entryNum < 1 {
      io.WriteString(w, "Journal entry number is invalid.")
      return
    }

    if journal_url == "entry" || journal_url == "title" {
      title := false

      if journal_url == "title" {
        title = true
      }

      str, err := readJournal(path, entNum, title)

      lastEntryNum[path] = entNum
      if err != nil {
        io.WriteString(w, "Couldn't read journal entry.")
        return
      }
      io.WriteString(w, str)
      return
    }

    if journal_url == "entryedit" || journal_url == "titleedit" {
      str := r.PostFormValue("text")

      title := false


      if journal_url == "titleedit" {
        title = true
      }

      if len(str) > 5000 && !title {
        io.WriteString(w, "Journal entry is too long!")
        return
      }

      if len(str) > 50 && title {
        io.WriteString(w, "Journal entry title is too long!")
        return
      }

      changeJournal(path, entNum, str, title)
      return
    }

    // NOTE: I don't think you can get to this point, but I might be wrong
    fmt.Println("Reached the end of entry handler without returning (???)")
    io.WriteString(w, "Error!")
    return
  }

  if journal_url == "last.js" {
    reqInt, err := strconv.Atoi(lastEntryNum[path])
    if err != nil || reqInt < 1 || reqInt > 10000 {
      lastEntryNum[path] = "1"
    }
    io.WriteString(w, "var reqNumber = " + lastEntryNum[path] + ";")
    return
  }

  if journal_url == "theme.js" {
    dat, err := readTheme(path)

    if err != nil {
      fmt.Println("Couldn't read a users theme")
    }

    io.WriteString(w, "var theme = '" + string(dat) +"';")
    return

  }
  if journal_url == "theme.css" {
    var theme string
    thm, err := readTheme(path)

    if err != nil {
        fmt.Println("Something went wrong while reading the theme")
    }

    switch string(thm) {
    case "normal":
      theme = "normal"
    case "dark":
      theme = "dark"
    case "darkblue":
      theme = "darkblue"
    case "red":
      theme = "red"
    case "grey":
      theme = "grey"
    default:
      theme = "normal"
    }

    dat, err := os.Open("./themes/" + theme + ".css")
    if err != nil {
      io.WriteString(w, "/* Theme file not found */")
      fmt.Println("A user requested a theme that wasn't avalible (" + theme + ")")
      return
    }
    w.Header().Add("Content-Type", "text/css")
    io.Copy(w, dat)
    return
  }

  // Handle settings being changed (Should happen every time the body is clicked)
  if journal_url == "settingschange" {
    theme := r.PostFormValue("theme")

    if len(theme) > 20 {
      fmt.Println("A user tried setting a theme that was over 20 characters long!")
      io.WriteString(w, "The theme is invalid!")
      return
    }

    err := setTheme(path, theme)

    if err != nil {
      io.WriteString(w, "Error")
      fmt.Println("Couldn't change theme for user: " + path)
      return
    }
    io.WriteString(w, "Settings changed")
    return
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
