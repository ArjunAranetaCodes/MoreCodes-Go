//Problem 7: Write a program that counts vowels in a String using Regular Expression.
package main

import "fmt"
import "regexp"

func main() {
  re := regexp.MustCompile(`[^aeiouAEIOU]`)
  word := re.ReplaceAllString(`Hello World`, ``)
  fmt.Println(len(word))
}
