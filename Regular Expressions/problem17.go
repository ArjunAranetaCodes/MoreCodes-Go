//Problem 17: Write a program that removes the last word in a string using Regular Expression.
package main

import "fmt"
import "regexp"

func main() {
  re := regexp.MustCompile(`\w+$`)
  word := re.ReplaceAllString(`Hello World`, ``)
  fmt.Println(word)
}
