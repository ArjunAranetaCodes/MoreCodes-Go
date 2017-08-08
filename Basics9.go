//Basics of Functions

package main

import "fmt"

func function1(){
 fmt.Println("Hello there")
}

func function2(num int){
 fmt.Println("That number is", num)
}

func function3() int{
 sum := 1 + 1
 return sum
}

func function4(firstName string, lastName string) string{
 fullname := firstName + " " + lastName;
 return fullname;
}

func main(){
 function1();
 function2(5);
 fmt.Println("It's true! 1 + 1 =", function3())
 fmt.Println("Hi", function4("More", "Codes"))
}



