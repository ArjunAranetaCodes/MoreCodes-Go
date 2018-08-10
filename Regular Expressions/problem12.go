//Problem 12: Write a program that counts the occurrence of digits in a string using Regular Expression.
package main

import "fmt"
import "regexp"

func main() {
  re := regexp.MustCompile(`\D`)
  word := re.ReplaceAllString(`Hello12 World12`, ``)
  fmt.Println(len(word))
}
