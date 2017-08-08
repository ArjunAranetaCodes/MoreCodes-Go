//Variables and Input

package main

import "fmt"

func main(){
 name := ""
 sex := ""
 age := 0

 fmt.Print("What is your name? : ")
 fmt.Scan(&name)
 fmt.Print("What is your sex? (M or F) ")
 fmt.Scan(&sex)
 fmt.Print("What is your age? ")
 fmt.Scan(&age)

 fmt.Println("Name: ", name)
 fmt.Println("Sex: ", sex)
 fmt.Println("Age: ", age)
}



