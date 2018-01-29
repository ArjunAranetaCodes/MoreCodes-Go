//Problem 9: Write a program that converys a hexadecimal number to decimal number.
package main

import "fmt"
import "strings"
import "math"

func Reverse(s string) string {
 runes := []rune(s)
 for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
  runes[i], runes[j] = runes[j], runes[i]
 }
 return string(runes)
}

func main(){
 dec := 0
 hex := "100"
 reverse  := Reverse(hex)
 count := 1

 for _, r := range reverse {
  c := string(r)
  if(strings.Contains("1", c)){
   dec = dec + int(math.Pow(float64(16), float64(count)))
  }
  count = count + 1
 }
 
 fmt.Print(dec)
}