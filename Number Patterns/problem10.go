/*
Problem 10: Write a program to print the number pattern using loop.
12345
23456
34567
45678
56789
*/

package main

import "fmt"

func main() {
	row := 5;
	col := 5;
	min := 1;
	for y := 0; y < row; y++ {
		num := min + y;
		for x := 0; x < col; x++ {
   fmt.Print(num);	
			num = num + 1;
		}
  fmt.Print("\n");	
	}
}
