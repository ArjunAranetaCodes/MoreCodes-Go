/*
Problem 16: Write a program to print the number pattern using loop.
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
	for y := row; y > 0; y-- {
		for x := 0; x < y; x++ {
			fmt.Print(x + 1);
		}
		fmt.Println();
	}
}
