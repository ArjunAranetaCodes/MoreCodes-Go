
package main

import "fmt"

func main() {
 num := 5
 fact := 1

 for x := num; x >= 1; x-- {
  fact = fact * x
 }
 fmt.Println("Factorial of 5 is ", fact)
}


