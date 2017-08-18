
package main

import "fmt"

func main() {
 array1 := [3] int{1,2,3}
 var array2 [3]int
 y := 0

 for x := len(array1) - 1; x >= 0; x-- {
  array2[y] = array1[x]
  y++
 }

 for _, x := range array2 {
  fmt.Print(x)
 }
}

