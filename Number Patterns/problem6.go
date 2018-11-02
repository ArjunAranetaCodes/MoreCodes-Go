/*
Problem 6: Write a program to print the number pattern of ones and zeros using loop.
11111
10001
10001
10001
11111
*/

package main

import "fmt"

func main() {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if((y == 0) || (y == 4) || (x == 0) || (x == 4)){
    fmt.Print("1");	
			}else{
    fmt.Print("0");	
			}			
		}
  fmt.Print("\n");	
	}
}
