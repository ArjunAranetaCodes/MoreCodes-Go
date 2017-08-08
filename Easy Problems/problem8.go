
package main

import "fmt"

func main(){
 sum := 0
 for x := 0; x <= 30; x++{
  if x % 5 == 0{
   sum = sum + x;
  }
 }

 fmt.Print("Sum of numbers divisible by 5 of 1 to 30 is ", sum)

}


