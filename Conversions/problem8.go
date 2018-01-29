//Problem 8: Write a program that converts a decimal number to hexadecimal number.
package main

import (
 "strconv"
 "fmt"
)

func main(){
 dec := 256
 hex  := ""
 
 for dec > 0{
  hex  = strconv.Itoa(dec % 16) + hex 
  dec = dec / 16
 }
 
 fmt.Print(hex)
}