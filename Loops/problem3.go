
package main

import "fmt"

func main() {
 word := "MoreCodes"
 letter := "a"
 letterCount := 0

 for _, r := range word {
  letter = string(r)
  fmt.Print(letter)
  letterCount = letterCount + 1
 }

 fmt.Println("\nTotal: ", letterCount)
}


