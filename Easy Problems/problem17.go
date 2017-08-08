
package main

import "fmt"
import "strings"

func main(){
 word := "program"
 letter := "a"

 if strings.Contains(word, letter) {
  fmt.Print("Contains a")
 }else{
  fmt.Print("Does not contain a")
 }
}


