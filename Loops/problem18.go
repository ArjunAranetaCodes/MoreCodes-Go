
package main

import "fmt"

func main() {
 for x := 0; x <= 11; x++ {
  fmt.Println(Fibonacci(x))
 }
}

func Fibonacci(num int) int{
 if (num == 1) || (num == 0){
  return num
 }else{
  return Fibonacci(num - 1) + Fibonacci(num - 2)
 }
}


