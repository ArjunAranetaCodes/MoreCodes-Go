//Problem 5: Write a program that matches time in 24 hour format.
package main

import "fmt"
import "regexp"

func main() {
    match1, _ := regexp.MatchString("^(0?[1-9]|1[012])(:[0-5]\\d) [APap][mM]$", "13:00")
    fmt.Println(match1)
    match2, _ := regexp.MatchString("^(0?[1-9]|1[012])(:[0-5]\\d) [APap][mM]$", "8:01 am")
    fmt.Println(match2)
}
