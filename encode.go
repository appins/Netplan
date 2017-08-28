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
  Titles []string
  Theme string
}

// This is used to change a journal entry, It can also be used to write titles
func changeJournal(journalName string, entry string, text string, title bool) error {
  dat, err := getJournalRaw(journalName)
  if err != nil {
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
  // too many entries or titles

  length := i + 1

  if title {

    if len(j.Titles) > length {
      length = len(j.Titles)
    }
    titlearr := make([]string, length)
    copy(titlearr, j.Titles)
    titlearr[i] = text
    j.Titles = titlearr

    jsonvalue, _ := json.Marshal(j)
    writeJournalRaw(journalName, jsonvalue)

    return nil
  }

  if len(j.Entries) > length {
    length = len(j.Entries)
  }
  entryarr := make([]string, length)
  copy(entryarr, j.Entries)
  entryarr[i] = text
  j.Entries = entryarr

  jsonvalue, _ := json.Marshal(j)
  writeJournalRaw(journalName, jsonvalue)

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
  j := Journal{arr, arr, "normal"}
  jsondata, _ := json.Marshal(j)

  ioutil.WriteFile("./entries/" + jorunalName + ".json", jsondata, 0777)

  return nil
}

// Read an entry from a journal (Can also read titles if the title bit is set)
func readJournal(journalName string, entryNumber string, title bool) (string, error) {
  dat, err := getJournalRaw(journalName)

  if err != nil {
    fmt.Println("Couldn't open a journal for reading")
    return "Couldn't open journal", err
  }

  var j Journal
  err = json.Unmarshal(dat, &j)

  if err != nil {
    fmt.Println("Couldn't convert JSON")
    return "Couldn't convert json", err
  }

  i, err := strconv.Atoi(entryNumber)
  i--

  if err != nil {
    fmt.Println("Entry number was invalid")
    return "Entry number invalid", err
  }

  if i > 2499 {
    return "Error reading journal entry", nil
  }

// Send the user a generic title if the one they requested is blank
  if len(j.Titles) - 1 < i && title {
    return "Entry title", nil
  }

  // Send the user a generic entry if the one they requested is blank
  if len(j.Entries) - 1 < i && !title {
    if i == 0 {
      // However, if it's also the first of the journla, read from a file.
      // (This file _SHOULD_ contain a welcome message)
      dat, _ := ioutil.ReadFile("./public/introtext.txt")
      return string(dat), nil
    }
    return "New planner entry", nil
  }

  if title {
    return j.Titles[i], nil
  }

  return j.Entries[i], nil
}

// Read a journal's theme really easily
func readTheme(journalName string) (string, error) {
  dat, err := getJournalRaw(journalName)

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
  dat, err := getJournalRaw(journalName)

  if err != nil {
    fmt.Println("Couldn't read a json file")
    return err
  }

  // Unmarshal the JSON
  var j Journal
  err = json.Unmarshal(dat, &j)

  if err != nil {
    fmt.Println("Couldn't unmarshal data (" + journalName + ".json)")
  }

  j.Theme = theme

  jsondata, _ := json.Marshal(j)

  // Errors aren't handled here because we were able to read from the same file.
  // If some error occurs, the theme won't be set.
  ioutil.WriteFile("./entries/" + journalName + ".json", jsondata, 0777)

  return nil
}
