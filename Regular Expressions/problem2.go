//Problem 2: Write a program that matches email address.
package main

import "fmt"
import "regexp"

func main() {
    match1, _ := regexp.MatchString("^[_A-Za-z0-9-\\+]+(\\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\\.[A-Za-z0-9]+)*(\\.[A-Za-z]{2,})$", "mark@yahoo.com")
    fmt.Println(match1)
    match2, _ := regexp.MatchString("^[_A-Za-z0-9-\\+]+(\\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\\.[A-Za-z0-9]+)*(\\.[A-Za-z]{2,})$", "mark-yahoo.com")
    fmt.Println(match2)
}
