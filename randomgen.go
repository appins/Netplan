package main

import (
  "crypto/rand"
)

// Get a random string from an array of strings
func randFromArr(arr []string) string {
  var x string
  for {
    oneB := make([]byte, 1)
    rand.Read(oneB)
    if int(oneB[0]) < len(arr) {
      x = arr[oneB[0]]
      break
    }
  }

  return x
}

// Generate a word type thing
func getRandom() string {
  vows := []string {"a", "e", "i", "o", "ee", "ea", "oo", "y", "ai"}
  cons := []string {"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p",
  "r", "s", "t", "v", "w", "x", "ch", "ll", "sh", "sh", "th"}
  nums := []string {"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "-"}

  build := ""

  build += randFromArr(cons)
  build += randFromArr(vows)
  build += randFromArr(cons)
  build += randFromArr(vows)
  build += randFromArr(cons)
  build += randFromArr(vows)
  build += randFromArr(nums)
  build += randFromArr(nums)
  build += randFromArr(nums)
  build += randFromArr(nums)

  return build
}
