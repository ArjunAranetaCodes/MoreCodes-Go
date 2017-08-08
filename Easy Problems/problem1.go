
package main

import "fmt"

func main(){
 num1 := 0
 num2 := 0

 fmt.Print("Enter value of num1: ")
 fmt.Scan(&num1)
 fmt.Print("Enter value of num2: ")
 fmt.Scan(&num2)
 sum := num1 + num2

 fmt.Println("Sum is", sum)
}


