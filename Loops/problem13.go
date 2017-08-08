
package main

import "fmt"

func main() {
 word := "MoreCodes"
 newWord := ""
 for x := len(word) - 1; x >= 0; x-- {
  newWord = newWord + string(word[x])
 }

 fmt.Print(newWord)
}



