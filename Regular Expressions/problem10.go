//Problem 10: Write a program that prints an integer with commas separator using Regular Expression.
package main

import "fmt"
import "regexp"

func main() {
  re := regexp.MustCompile(`(\d)(\?=(\d{3})+$)`)
  word := re.ReplaceAllString(`1000`, `1,`)
  fmt.Println(word)
}
