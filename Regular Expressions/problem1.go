//Problem 1: Write a program to test if the first character of the String is uppercase.
package main

import "fmt"
import "regexp"

func main() {
    match1, _ := regexp.MatchString("^[A-Z][a-z0-9_-]{3,19}$", "Hello")
    fmt.Println(match1)
    match2, _ := regexp.MatchString("^[A-Z][a-z0-9_-]{3,19}$", "world")
    fmt.Println(match2)
}
