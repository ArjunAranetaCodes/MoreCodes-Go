
package main

import "fmt"

func main() {
 num1 := 3
 num2 := 5
 prod := 0
 for x := 1; x <= num2; x++ {
  prod = prod + num1
 }

 fmt.Print("3 * 5 = ", prod)
}


