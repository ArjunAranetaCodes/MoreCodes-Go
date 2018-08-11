//Problem 18: Write a program that extracts string inside quotation marks using Regular Expression.
package main

import "fmt"
import "regexp"

func main() {
    r, _ := regexp.Compile("\\'([^\"]*)\\'")
    fmt.Println(r.FindString("Hello 'World'"))
}
