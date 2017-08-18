
package main

import "fmt"

func main() {
 arrNumbers := [3] int{1,2,3}
 indexOfNum := 0
 x := 0

 for x = 0; x < len(arrNumbers); x++ {
  if arrNumbers[x] == 2{
   indexOfNum = x
  }
 }
 fmt.Println(indexOfNum)
}

