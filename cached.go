package main

import (
  "fmt"
  "time"
  "io/ioutil"
  "os"
)

var lock bool
var clearLock bool

// This can be called to open a journal to one of the maps
// This also checks if the journal is already open, so it should be called
// every time. It returns true if the journal exists
func openJournal(journal string) error {
  // Make sure journals aren't being cleared
  for clearLock {}

  // This keeps this function from interfering with the clearning of journals
  lock = true

  if openJournals[journal] != nil {
    lock = false
    return nil
  }

  dat, err := ioutil.ReadFile("./entries/" + journal + ".json")
  if err != nil {
    fmt.Println("Couldn't open requested journal (openJournal(" + journal + "))")
    lock = false
    return err
  }

  openJournals[journal] = dat
  fmt.Println("Opened a journal")

  lock = false
  return nil
}

// This is used to return if a journals content exists. It uses caching.
func getJournalRaw(journal string) ([]byte, error) {
  err := openJournal(journal)
  if err == nil {
    return openJournals[journal], nil
  }
  return []byte(""), err
}

func writeJournalRaw(journal string, data []byte) error {
  err := openJournal(journal)
  if err == nil {
    openJournals[journal] = data
    return nil
  }
  return err
}

// This function (on a seperate thread) will clear the array of open jorunals
// and also write all of them to your hard drive
func cacheClearAndWrite () {
  for {
    // This waits 100 seconds, this should be changed to a larger value
    // depending on use. If the period is too long, the memory might fill up
    // and the program might crash. Too short would be inefficent.
    time.Sleep(10 * time.Second)
    fmt.Println("Writing journals")
    clearLock = true
    for {
      if !lock {
        break
      }
      time.Sleep(time.Second)
    }

    for k, v := range openJournals {
      ioutil.WriteFile("./entries/" + k + ".json", v, 0777)
      delete(openJournals, k)
    }

    fmt.Println("Cleared and wrote journals")

    if pathExists("./CLOSE_NETPLAN") {
      fmt.Println("Killing Netplan")
      ioutil.WriteFile("KILL_NETPLAN_NOW", []byte(""), 0777)
      os.Exit(0)
    }
    clearLock = false
  }
}
