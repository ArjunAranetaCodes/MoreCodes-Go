//Problem 15: Write a program that takes a value inside a <div> tag using Regular Expression.
package main

import "fmt"
import "regexp"

func main() {
  re := regexp.MustCompile(`<(\"[^\"]*\"|\'[^\']*\'|[^\'\">])*>`)
  word := re.ReplaceAllString(`<div>Hello World</div>`, ``)
  fmt.Println(word)
}
