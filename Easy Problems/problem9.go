
package main

import "fmt"

func main(){
 sum := 0
  num := 0

  fmt.Print("Enter value of num: ")
  fmt.Scan(&num)

 for x := 0; x <= num; x++{
  sum = sum + x;
 }

 fmt.Print("Sum of 1 to ", num, " is ", sum)

}


