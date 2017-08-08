
package main

import "fmt"

func main() {
 num := 371
 sum := 0
 temp := 0
 rmdr := 0

 temp = num

 for temp != 0{
  rmdr = temp % 10
  sum = sum + (rmdr * rmdr * rmdr)
  temp = temp / 10
 }

 if num == sum{
  fmt.Println("Armstrong number");
 }else {
  fmt.Println("Not an Armstrong number");
 }
}


