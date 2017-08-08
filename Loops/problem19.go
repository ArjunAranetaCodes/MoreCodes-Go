
package main

import "fmt"

func main() {
 PrintEven(10);
}

func PrintEven(num int) int{
 if (num == 0){
  return num
 }else{
  if (num % 2 == 0){
   fmt.Println(num)
  }
  return PrintEven(num - 1)
 }
}


