
package main

import "fmt"

func main() {
 array1 := [3] int{1,2,3}
 var array2 [3]int

 array2 = array1

 for _, x := range array2 {
  fmt.Println(x)
 }
}

