/*
Problem 1: Write a program to print the number pattern of ones and zeros using loop.
11111
11111
11111
11111
11111
*/

package main

import "fmt"

func main() {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			fmt.Print("1");
		}
		fmt.Println();
	}
}
