/*
Problem 7: Write a program to print the number pattern of ones and zeros using loop.
10101
01010
10101
01010
10101
*/

package main

import "fmt"

func main() {
	count := 0;
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			count += 1;
			if(count % 2 == 1){
    fmt.Print("1");	
			}else{
    fmt.Print("0");	
			}			
		}
  fmt.Print("\n");	
	}
}
