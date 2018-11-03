/*
Problem 15: Write a program to print the number pattern using loop.
11111
2222
333
44
5
*/

package main

import "fmt"

func main() {
	row := 5;
	num := 1;
	for y := row; y > 0; y-- {
		for x := 0; x < y; x++ {
			fmt.Print(num);	
		}
		num = num + 1;
		fmt.Print("\n");	
	}
}
