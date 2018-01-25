//Problem 4: Write a program that converts a string to array/list.
package main

import (
    "fmt"
    "encoding/json"
)

func main() {
    numbers  := "[2,15,23]"
    var arrNumbers []int
    json.Unmarshal([]byte(numbers), &arrNumbers)
    
    fmt.Print(arrNumbers)
}