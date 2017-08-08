
package main

import "fmt"

func main() {

 for y := 0; y <= 5; y++ {
  _ = y
  for x := 0; x < y; x++ {
   fmt.Print("*")
  }
  fmt.Println()
 }
}


