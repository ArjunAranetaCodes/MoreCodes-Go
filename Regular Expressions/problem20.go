//Problem 20: Write a program that validates if string length is between 5 to 10 using Regular Expression.
package main

import "fmt"
import "regexp"

func main() {
    match1, _ := regexp.MatchString("\\w{5,10}\\b", "Hello")
    fmt.Println(match1)
    match2, _ := regexp.MatchString("\\w{5,10}\\b", "Hi")
    fmt.Println(match2)
}
