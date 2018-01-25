//Problem 2: Write a program that converts a string to integer.
package main

import (
 "strconv"
 "fmt"
)

func main() {
 strnum := "10"
 num, err := strconv.ParseInt(strnum, 10, 64)
 _ = err
 fmt.Println(num)
}
