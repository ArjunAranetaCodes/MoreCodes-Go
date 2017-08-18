
package main

import "fmt"

func main() {
 arrNumbers := []int{1,2,3}
 indexOfNum := 0

 for x := 0; x < len(arrNumbers); x++ {
  if arrNumbers[x] == 2{
   indexOfNum = x
  }
 }

 i := indexOfNum
 arrNumbers[i] = arrNumbers[len(arrNumbers)-1]
 arrNumbers = arrNumbers[:len(arrNumbers)-1]

 for _, x := range arrNumbers {
  fmt.Println(x)
 }
}

