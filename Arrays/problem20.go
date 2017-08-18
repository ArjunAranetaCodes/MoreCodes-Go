
package main

import "fmt"

func main() {
 arrNumbers := []int{4,3,2,1}
 closest := 0
 numDiff := arrNumbers[0]

 for _ , x := range arrNumbers {
  diff := 0 - x
  if diff < 0 {
   diff = -diff
  }
  if(diff < numDiff){
   numDiff = diff
   closest = x
  }
 }

 fmt.Println("Closest to 0 is", closest)
}

