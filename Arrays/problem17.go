
package main

import "fmt"

func main() {
 arrNumbers := []int{1,2,3}
 arrNumbers = append(arrNumbers, 5)

 for _, x := range arrNumbers {
  fmt.Println(x)
 }
}

