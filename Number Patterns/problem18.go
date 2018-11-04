/*
Problem 18: Write a program to print the number pattern using loop.
  1  
 111
11111
 111
  1
*/

package main

import "fmt"

func main() {
 rows := 3;
 ones := 1; 
 blank := rows - 1;
   
 for y := 1; y < rows*2; y++ {
  for x := 1; x <= blank; x++ {
   fmt.Print(" ");
  }
  
  for x := 1; x<ones*2; x++ {
   fmt.Print("*");
  }
  
  fmt.Println();
  
  if(y<rows){
   blank--;
   ones++;
  }else{
   blank++;
   ones--;
  }
 }
}
