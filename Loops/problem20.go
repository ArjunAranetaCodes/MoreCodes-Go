
package main

import "fmt"

func main() {
 fmt.Println("Sum is ", GetSum(10, 0))
}

func GetSum(num int, sum int) int{
 if num == 0{
  return sum
 }else{
  return GetSum((num - 1), (sum + num))
 }
}


