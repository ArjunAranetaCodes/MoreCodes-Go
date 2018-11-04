/*
Problem 19: Write a program to print the number pattern using loop.
11 11
11 11
   
11 11
11 11
*/

package main

import "fmt"

func main() {
	row := 5;
	col := 5;
	for y := 0; y < row; y++ {
		for x := 0; x < col; x++ {
   if(col / 2 == x || row / 2 == y){
    fmt.Print(" ");
   }else if((col % 2 == 0 && (col / 2) == x) || (row % 2 == 0 && (row / 2) == y)){
    fmt.Print(" ");
   }else{
    fmt.Print("1");
   }		
		}
		fmt.Print("\n");
	}
}
