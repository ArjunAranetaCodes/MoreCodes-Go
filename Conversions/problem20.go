//Problem 20: Write a program that converts a number to its 
//corresponding month (e.g. 1 = January).
package main

import (
 "fmt"
)

func main(){
  index := 1
  month := []string{"January", "February", "March", "April",
   "May", "June", "July", "August",
   "September", "October", "November", "December"}     
    
  fmt.Print(month[index - 1])
}