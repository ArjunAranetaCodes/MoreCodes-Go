//Problem 10: Write a program that converts a decimal number to octal number.
package main

import (
 "strconv"
 "fmt"
)

func main(){
 dec := 256
 oct  := ""
 
 for dec > 0{
  oct  = strconv.Itoa(dec % 8) + oct 
  dec = dec / 8
 }
 
 fmt.Print(oct)
}