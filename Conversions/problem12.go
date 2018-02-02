//Problem 12: Write a program that converts numbers of day to seconds.
package main

import (
 "fmt"
)

func main(){
 day := 1
 seconds := day * 24 * 60 * 60
 
 fmt.Print(seconds)
}