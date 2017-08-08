
package main

import "fmt"
import "strings"

func main(){
 word1 := "programming"
 word2 := "program"

 if strings.Contains(word1, word2) {
  fmt.Print(word2, " found")
 }else{
  fmt.Print(word2, " not found")
 }
}


