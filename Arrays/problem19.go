
package main

import "fmt"

func main() {
 array1 := []int{1,2,3}
 array2 := []int{1,2,3}
 array1 = append(array1, array2...)

 for _, x := range array1 {
  fmt.Println(x)
 }
}

