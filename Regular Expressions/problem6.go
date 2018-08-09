//Problem 6: Write a program that removes white space and non-visible characters.
package main

import "fmt"
import "regexp"

func main() {
  re := regexp.MustCompile(`\s`)
  word := re.ReplaceAllString(`Hello World`, ``)
  fmt.Println(word)
}
