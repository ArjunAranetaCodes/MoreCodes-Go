
package main

import "fmt"

func main() {
 num := 3333
 sum := 0
 temp := 0
 rmdr := 0

 temp = num

 for temp != 0{
  rmdr = temp % 10
  sum = sum * 10 + rmdr
  temp = temp / 10
 }

 if num == sum{
  fmt.Println("Palindrome number")
 }else {
  fmt.Println("Not an Palindrome number")
 }
}


