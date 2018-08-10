//Problem 11: Write a program that counts the occurrence of a string in a string using Regular Expression.
package main

import "fmt"
import "regexp"

func main() {
  word := `HelloWorldHelloWorld`;
  pattern := `Hello`;
  re := regexp.MustCompile(pattern)
  count1 := len(word);
  count2 := len(re.ReplaceAllString(word, ``))
  fmt.Println((count1 - count2) / len(pattern))
}
