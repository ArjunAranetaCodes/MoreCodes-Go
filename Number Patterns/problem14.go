/*
Problem 14: Write a program to print the number pattern using loop.
1
22
333
4444
55555
*/

package main

import "fmt"

func main() {
	row := 5;
	for y := 0; y <= row; y++ {
		for x := 0; x < y; x++ {
			fmt.Print(y);	
		}
		fmt.Print("\n");	
	}
}
