
package main

import "fmt"

func main(){
 celsius := 0.0
 farenheit := 0.0

 fmt.Print("Enter value of celsius: ")
 fmt.Scan(&celsius)

 farenheit = (9.0/5.0) * celsius + 32.0

 fmt.Println(celsius, "C equals to ", farenheit, "F")
}


