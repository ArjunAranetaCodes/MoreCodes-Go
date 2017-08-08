
package main

import "fmt"

func main(){
 radius := 0

 fmt.Print("Enter value of radius: ")
 fmt.Scan(&radius)

 pi := 3.14
 dia := radius * radius
 area := pi * float64(dia)
 cir := 2.0 * pi * float64(radius)

 fmt.Println("Diameter of the circle is", dia)
 fmt.Println("Area of the circle is", area)
 fmt.Println("Circumference of the circle is", cir)
}


