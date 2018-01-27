//Problem 6: Write a program that converts a binary number to decimal number.
package main

import "fmt"
import "strings"

func Reverse(s string) string {
 runes := []rune(s)
 for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
  runes[i], runes[j] = runes[j], runes[i]
 }
 return string(runes)
}

func main(){
 dec := 0
 binary := "110"
 reverse  := Reverse(binary)
 count := 0

 for _, r := range reverse {
  c := string(r)
  if(strings.Contains("1", c)){
   dec = dec + (2^count)
  }
  count = count + 1
 }
 
 fmt.Print(dec)
}