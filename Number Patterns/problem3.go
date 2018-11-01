/*
Problem 3: Write a program to print the number pattern of ones and zeros using loop.
01010
01010
01010
01010
01010
*/

package main

import "fmt"

func main() {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if(x % 2 == 0){
    fmt.Print("0");	
			}else{
				fmt.Print("1");	
			}			
		}
		fmt.Print("\n");	
	}
}
