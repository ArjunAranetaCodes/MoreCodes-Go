//Problem 19: Write a program that counts all numbers in a string using Regular Expression.
package main

import "fmt"
import "regexp"

func main() {
  re := regexp.MustCompile(`[^0-9]`)
  word := re.ReplaceAllString(`Hello World123`, ``)
  fmt.Println(len(word))
}
