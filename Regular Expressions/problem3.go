//Problem 3: Write a program that checks if string contains sample format date of (yyyy-mm-dd)
package main

import "fmt"
import "regexp"

func main() {
    match1, _ := regexp.MatchString("([0-9]{4})-([0-9]{2})-([0-9]{2})", "2018-01-02")
    fmt.Println(match1)
    match2, _ := regexp.MatchString("([0-9]{4})-([0-9]{2})-([0-9]{2})", "01-01-02")
    fmt.Println(match2)
}
