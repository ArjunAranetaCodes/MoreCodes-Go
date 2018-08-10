//Problem 14: Write a program that recognizes valid hex color value using Regular Expression.
package main

import "fmt"
import "regexp"

func main() {
    match1, _ := regexp.MatchString("^([01]?\\d\\d?|2[0-4]\\d|25[0-5])\\.([01]?\\d\\d?|2[0-4]\\d|25[0-5])\\.([01]?\\d\\d?|2[0-4]\\d|25[0-5])\\.([01]?\\d\\d?|2[0-4]\\d|25[0-5])$", "192.168.1.1")
    fmt.Println(match1)
    match2, _ := regexp.MatchString("^([01]?\\d\\d?|2[0-4]\\d|25[0-5])\\.([01]?\\d\\d?|2[0-4]\\d|25[0-5])\\.([01]?\\d\\d?|2[0-4]\\d|25[0-5])\\.([01]?\\d\\d?|2[0-4]\\d|25[0-5])$", "1.1.1.1.1")
    fmt.Println(match2)
}
