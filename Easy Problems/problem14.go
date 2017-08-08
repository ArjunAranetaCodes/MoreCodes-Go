
package main

import "fmt"
import "strings"

func main(){
 word := "program"
 newWord := strings.Replace(word, "a", "e", -1)

 fmt.Print(newWord)
}


