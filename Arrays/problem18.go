
package main

import (
 "fmt"
 "sort"
)

func main() {
 arrNumbers := []int{1, 3, 2, 4}
 sort.Ints(arrNumbers)

 for _, x := range arrNumbers {
  fmt.Println(x)
 }
}

