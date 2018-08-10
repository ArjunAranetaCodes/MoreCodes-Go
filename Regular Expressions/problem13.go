//Problem 13: Write a program that recognizes valid hex color value using Regular Expression.
package main

import "fmt"
import "regexp"

func main() {
    match1, _ := regexp.MatchString("^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$", "#fff")
    fmt.Println(match1)
    match2, _ := regexp.MatchString("^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$", "#asdf")
    fmt.Println(match2)
}
