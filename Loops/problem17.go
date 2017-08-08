
package main

import "fmt"

func main() {
 num1 := 1
 num2 := 1

 fmt.Println(num1)
 for num2 < 100{
  fmt.Println(num2)
  num2 = num2 + num1
  num1 = num2 - num1
 }
}


