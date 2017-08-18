
package main

import "fmt"

func main() {
 arrNumbers := [4] int{1,1,2,3}
 count := 0

 for _, x := range arrNumbers {
  if x == 1{
   count = count + 1
  }
 }
 fmt.Println("Occurence :", count)
}

