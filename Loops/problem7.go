
package main

import "fmt"

func main() {
 n := 10
 sum := 0
 for x := 0; x <= n; x++ {
  sum = sum + x
 }

 fmt.Print("Sum is ", sum)
}


