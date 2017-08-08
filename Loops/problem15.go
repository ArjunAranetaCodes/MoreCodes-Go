
package main

import "fmt"

func main() {
 num := 0

 for num > -1{
  fmt.Print("Enter a number: ")
  fmt.Scan(&num)
 }

 fmt.Println("Terminated")
}


