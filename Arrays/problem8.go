
package main

import "fmt"

func main() {
 arrNumbers := [3] int{1,2,3}
 min := arrNumbers[0]

 for _, x := range arrNumbers {
  if x < min{
   min = x
  }
 }

 fmt.Println("Lowest Number: ", min)
}

