
package main

import "fmt"

func main(){
 num := 0

 fmt.Print("Enter value of num: ")
 fmt.Scan(&num)

 if num % 2 == 0 {
  fmt.Print("Number is even")
 }else{
  fmt.Print("Number is odd")
 }
}


