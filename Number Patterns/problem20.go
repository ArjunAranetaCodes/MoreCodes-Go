/*
Problem 20: Write a program to print the number pattern using loop.
  1
 222
33333
 444
  5
*/

package main

import "fmt"

func main() {
 rows := 3;
 nums := 1; 
 blank := rows - 1;
   
 for y := 1; y < rows*2; y++ {
  for x := 1; x <= blank; x++ {
   fmt.Print(" ");
  }
  
  for x := 1; x<nums*2; x++ {
   fmt.Print(y);
  }
  
  fmt.Println();
  
  if(y<rows){
   blank--;
   nums++;
  }else{
   blank++;
   nums--;
  }
 }
}
