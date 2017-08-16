package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
  "strconv"
)

// This is the format of a journal currently
type Journal struct {
  Entries []string
  Theme string
}

// This is used to change a journal entry
func changeJournal(journalName string, entry string, entrytext string) error {
  dat, err := ioutil.ReadFile("./entries/" + journalName + ".json")
  if err != nil {
    fmt.Println("Error reading journal")
    return err
  }

  var j Journal
  err = json.Unmarshal(dat, &j)

  if err != nil {
    fmt.Println("Error unmarshaling json")
    return err
  }

  i, err := strconv.Atoi(entry)
  i--

  if i > 2499 || i < 0 {
    fmt.Println("Wrote outside journal bounds")
    return nil
  }

  if err != nil {
    fmt.Println("Error converting number from string to integer")
    return err
  }

  // Make sure that the array does not access out of range but doesn't create
  // too many entries
  length := i + 1
  if len(j.Entries) > length {
    length = len(j.Entries)
  }
  entryarr := make([]string, length)
  copy(entryarr, j.Entries)
  entryarr[i] = entrytext
  j.Entries = entryarr

  jsonvalue, _ := json.Marshal(j)
  ioutil.WriteFile("./entries/" + journalName + ".json", jsonvalue, 0777)

  return nil
}

// This will create a new journal
func newJournal(jorunalName string) error {
  f, err := os.Create("./entries/" + jorunalName + ".json")
  defer f.Close()

  if err != nil {
    fmt.Println("error creating file (" + jorunalName + ".json)")
    return err;
  }

  var arr []string
  j := Journal{arr, "normal"}
  jsondata, _ := json.Marshal(j)

  ioutil.WriteFile("./entries/" + jorunalName + ".json", jsondata, 0777)

  return nil
}

// Read an entry from a journal
func readJournal(jorunalName string, entryNumber string) (string, error) {
  dat, err := ioutil.ReadFile("./entries/" + jorunalName + ".json")

  if err != nil {
    fmt.Println("Couldn't open a journal for reading")
    return "", err
  }

  var j Journal
  err = json.Unmarshal(dat, &j)

  if err != nil {
    fmt.Println("Couldn't convert JSON")
    return "", err
  }

  i, err := strconv.Atoi(entryNumber)
  i--

  if err != nil {
    fmt.Println("Entry number was invalid")
    return "", err
  }

  if i > 2499 {
    return "Error reading journal entry", nil
  }
  if len(j.Entries) - 1 < i {
    if i == 0 {
      dat, _ := ioutil.ReadFile("./public/introtext.txt")
      return string(dat), nil
    }
    return "New planner entry", nil
  }

  return j.Entries[i], nil
}

// Read a journal's theme really easily
func readTheme(journalName string) (string, error) {
  dat, err := ioutil.ReadFile("./entries/" + journalName + ".json")

  if err != nil {
    fmt.Println("Couldn't read a json file")
    return "", err
  }

  var j Journal
  err = json.Unmarshal(dat, &j)

  if err != nil {
    fmt.Println("Couldn't unmarshal data (" + journalName + ".json)")
  }

  return j.Theme, nil
}

// Set a journal's theme
func setTheme(journalName string, theme string) error {
  dat, err := ioutil.ReadFile("./entries/" + journalName + ".json")

  if err != nil {
    fmt.Println("Couldn't read a json file")
    return err
  }

  var j Journal
  err = json.Unmarshal(dat, &j)

  if err != nil {
    fmt.Println("Couldn't unmarshal data (" + journalName + ".json)")
  }

  j.Theme = theme

  jsondata, _ := json.Marshal(j)

  ioutil.WriteFile("./entries/" + journalName + ".json", jsondata, 0777)

  return nil
}
