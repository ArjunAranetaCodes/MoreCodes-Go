
package main

import "fmt"

func main() {
 arrNumbers := [3] int{1,2,3}
 sum := 0

 for _, x := range arrNumbers {
  sum = sum + x
 }

 fmt.Println("Average is ", sum / 3)
}

