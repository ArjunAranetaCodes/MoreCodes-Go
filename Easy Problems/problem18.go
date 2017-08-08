
package main

import "fmt"
import "strings"

func main() {
 word := "program"
 letter := "a"
 letterCount := 0

 for _, r := range word {
  c := string(r)
  if(strings.Contains(letter, c)){
   letterCount = letterCount + 1
  }
 }

 fmt.Println("Total: ", letterCount)
}


