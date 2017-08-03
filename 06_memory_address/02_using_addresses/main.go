package main

import "fmt"

const metersToYard  float64 = 1.09361


// This function basically gets an input from the user and converts value based on given formula.
// A variable is declared initially, this means that it has a memmory address allocated.
// use the Scan function of fmt pacakage fmt.Scan to get the input.
// The user entrerd input needs to be stored somewhere. This is where we need to give the memory address where it needs to stored to
// Here we have decalared a variable meters of type float64.
// We want to save the user provided input to this variable, so we give address of meters as argument to the fmt.Scan

func main ()  {
  fmt.Printf("Printing Something\n")
  var meters float64
  fmt.Println("Enter the meter value to be converted to Yards ")
  fmt.Scan(&meters)
  yards := meters * metersToYard

  fmt.Println(" The meter value is ",meters, " The Yards value is ", yards)
}
