//Problem 16: Write a program that takes a value inside a <a> tag using Regular Expression.
package main

import "fmt"
import "regexp"

func main() {
  re := regexp.MustCompile(`<(\"[^\"]*\"|\'[^\']*\'|[^\'\">])*>`)
  word := re.ReplaceAllString(`<a>Hello World</a>`, ``)
  fmt.Println(word)
}
