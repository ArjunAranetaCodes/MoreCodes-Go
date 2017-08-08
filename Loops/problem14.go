
package main

import "fmt"

func main() {
 num := 0
 sum := 0
 ave := 0

 for x := 0; x < 5; x++{
  fmt.Print("Enter a number: ")
  fmt.Scan(&num)
  sum = sum + num
 }
 ave = sum / 5
 fmt.Println("Average is ", ave)
}


