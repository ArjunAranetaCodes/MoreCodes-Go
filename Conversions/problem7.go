//Problem 7: Write a program that converts a decimal number to binary number.
package main

import (
 "strconv"
 "fmt"
)

func main(){
 dec := 10
 binary := ""
 
 for dec > 0{
  binary = strconv.Itoa(dec % 2) + binary
  dec = dec / 2
 }
 
 fmt.Print(binary)
}