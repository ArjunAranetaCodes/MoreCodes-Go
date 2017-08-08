
package main

import "fmt"

func main(){
 num1 := 0
 num2 := 0
 num3 := 0
 ave := 0

 fmt.Print("Enter value of num1: ")
 fmt.Scan(&num1)
 fmt.Print("Enter value of num2: ")
 fmt.Scan(&num2)
 fmt.Print("Enter value of num3: ")
 fmt.Scan(&num3)
 ave = (num1 + num2 + num3) / 3
 fmt.Print("Average is ", ave)
}


