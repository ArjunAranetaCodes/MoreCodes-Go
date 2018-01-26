//Problem 5: Write a program that converts an array/list to string.
package main

import (
  "fmt"
  "strings"
)

func main() {
  arrNumbers := []string {"1","2","3"}
  numbers := strings.Join(arrNumbers," ")
  fmt.Println(numbers)
}