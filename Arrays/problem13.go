
package main

import "fmt"

func main() {
 arrNumbers := [4] int{1,2,3,4}

 for _, x := range arrNumbers {
  if x % 2 == 0{
   fmt.Println(x)
  }
 }
}

