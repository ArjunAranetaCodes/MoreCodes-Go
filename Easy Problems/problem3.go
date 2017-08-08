
package main

import "fmt"

func main(){
 length := 0
 width := 0

 fmt.Print("Enter value of length: ")
 fmt.Scan(&length)
 fmt.Print("Enter value of width: ")
 fmt.Scan(&width)
 area := length * width

 fmt.Println("Area of Rectangle is", area)
}


