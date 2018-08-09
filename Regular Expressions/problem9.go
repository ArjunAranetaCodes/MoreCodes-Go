//Problem 9: Write a program that checks if the string is alphanumeric using Regular Expression.
package main

import "fmt"
import "regexp"

func main() {
    match1, _ := regexp.MatchString("(([A-Z].*[0-9])|([0-9].*[A-Z]))", "HelloWorld")
    fmt.Println(match1)
    match2, _ := regexp.MatchString("(([A-Z].*[0-9])|([0-9].*[A-Z]))", "HelloWorld123")
    fmt.Println(match2)
}
