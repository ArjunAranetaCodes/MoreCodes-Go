//Problem 4: Write a program that matches time in 12 hour format.
package main

import "fmt"
import "regexp"

func main() {
    match1, _ := regexp.MatchString("(((0[1-9])|(1[0-2])):([0-5])([0-9])\\s(a|p)m)", "08:01 am")
    fmt.Println(match1)
    match2, _ := regexp.MatchString("(((0[1-9])|(1[0-2])):([0-5])([0-9])\\s(a|p)m)", "18:01 pm")
    fmt.Println(match2)
}
