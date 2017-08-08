
package main

import "fmt"
import "strings"

func main() {
 word := "program"
 vowels := "aeiou"
 vowelCount := 0

 for _, r := range word {
  c := string(r)
  if(strings.Contains(vowels, c)){
   vowelCount = vowelCount + 1
  }
 }

 fmt.Println("Total: ", vowelCount)
}


