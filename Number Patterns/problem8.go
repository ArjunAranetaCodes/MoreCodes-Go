/*
Problem 8: Write a program to print the number pattern of ones and zeros using loop.
11111
11111
11011
11111
11111
*/

package main

import "fmt"

func main() {
	row := 5;
	col := 5;
	for y := 0; y < row; y++ {
		for x := 0; x < col; x++ {
			if(x == (row / 2) && y == (col / 2)){
    fmt.Print("0");	
			}else{
    fmt.Print("1");	
			}			
		}
  fmt.Print("\n");	
	}
}
