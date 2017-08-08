//If, ElseIf, and Else Structure

package main

import "fmt"

func main(){
 num1 := 1
 num2 := 2

 if num1 > num2 {
  fmt.Println("First number is higher!")
 }

 if num1 == num2 {
  fmt.Println("They are equal!");
 }else{
  fmt.Println("They are not equal!");
 }

 if num1 > num2{
  fmt.Println("First number is greater!");
 }else if num1 < num2{
  fmt.Println("Second number is greater!");
 }
}



