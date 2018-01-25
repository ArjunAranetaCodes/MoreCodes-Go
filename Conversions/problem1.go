//Problem 1: Write a program that converts a number to string.
package main

import (
 "strconv"
 "fmt"
)

func main() {
 num := 10
 strnum := strconv.Itoa(num)
 fmt.Println(strnum)
}
