/*
Problem 17: Write a program to print the number pattern using loop.
1
12
123
1234
12345
1234
123
12
1
*/

package main

import "fmt"

func main() {
	row := 5;
	for y := 0; y <= row; y++ {
		for x := 0; x < y; x++ {
			fmt.Print(x + 1);
		}
		fmt.Println();
	}
	
	for y := row - 1; y > 0; y-- {
		for x := 0; x < y; x++ {
			fmt.Print(x + 1);
		}
		fmt.Println();
	}
}
