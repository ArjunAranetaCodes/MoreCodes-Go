
package main

import "fmt"

func main() {
 arrNumbers := [3] int{1,2,3}
 max := arrNumbers[0]

 for _, x := range arrNumbers {
  if x > max{
   max = x
  }
 }

 fmt.Println("Largest Number: ", max)
}

