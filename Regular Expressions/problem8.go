//Problem 8: Write a program that checks if the String if valid url using Regular Expression.
package main

import "fmt"
import "regexp"

func main() {
    match1, _ := regexp.MatchString("^(https?:\\/\\/)?(www\\.)?([\\w]+\\.)+[‌​\\w]{2,63}\\/?$", "http://www.example.com")
    fmt.Println(match1)
    match2, _ := regexp.MatchString("^(https?:\\/\\/)?(www\\.)?([\\w]+\\.)+[‌​\\w]{2,63}\\/?$", "wwwexamplecom")
    fmt.Println(match2)
}
