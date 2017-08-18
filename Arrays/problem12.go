
package main

import "fmt"

func main() {
 array1 := [3] int{1,2,3}
 array2 := [3] int{1,2,3}

 if array1 == array2 {
  fmt.Println("Equal Arrays")
 }else{
  fmt.Println("Not Equal")
 }
}

