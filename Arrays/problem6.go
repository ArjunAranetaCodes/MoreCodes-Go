
package main

import "fmt"

func main() {
 arrNumbers := [3] int{1,2,3}
 count := 0

 for _, x := range arrNumbers {
  if x == 2{
   count = count + 1
  }
 }

 if count > 0 {
  fmt.Println("Contains 2");
 }else{
  fmt.Println("Does not contain 2")
 }
}

