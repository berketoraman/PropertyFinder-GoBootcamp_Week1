//These are Chapter2 examples of the book Go Programming Language
package main

import (
"fmt" 
)

const boilingF = 212.0

type Celsius float64
type Fahrenheit float64

const (
     AbsoluteZeroC Celsius = -273.15
     FreezingC     Celsius = 0
     BoilingC      Celsius = 100
)

func main() {
   
   // Boiling prints the boiling point of water.
   var f = boilingF
   var c = (f - 32) * 5 / 9
   fmt.Printf("boiling point = %g°F or %g°C\n", f, c)// Output: boiling point = 212°F or 100°C

   // Ftoc prints two Fahrenheit-to-Celsius conversions.
   const freezingF, boilingF = 32.0, 212.0
   fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) //Output: "32°F = 0°C"
   fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   //Output: "212°F = 100°C"
 
   //Pointers
   x := 1
   p := &x         // p, of type *int, points to x
   fmt.Println(*p) // Output: "1"
   *p = 2          // equivalent to x = 2
   fmt.Println(x)  // Output: "2"

   //example of function incr(p *int) 
   v := 1
   incr(&v)              // side effect: v is now 2
   fmt.Println(incr(&v)) // "3" (and v is 3)

   //Greatest common divider and Fibonacci functions examples
   fmt.Println(gcd(10,20)) //10
   fmt.Println(gcd(36,240)) //12
   fmt.Println(fib(12)) //144
   fmt.Println(fib(7)) //13
	
   //Using the constant values to convert between Celsius and Fahrenheit
   fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
   fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
 
   //The program below has three different variables called z because each declaration appears in a different lexical block. (This example illustrates scope rules, not good style!)
   z := "hello!"
   for i := 0; i < len(z); i++ {
        z := z[i]
        if z != '!' {
            z := z + 'A' - 'a'
			fmt.Printf("%c\n", z) // "HELLO" (one letter per iteration)
			} 
	}

	//The example below also has three variables named a, each declared in a different block.one in the function body, one in the for statement’s block, and one in the loop body but only two of the blocks are explicit:
    a:= "hello"
    for _, a := range a {
        a := a + 'A' - 'a'
        fmt.Printf("%c\n", a) // "HELLO" (one letter per iteration)
	}
}

//function for Fahrenheit-to-Celsius conversion
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

//function to increase the value of the pointer by one.
func incr(p *int) int {
 *p++ // increments what p points to; does not change p
 return *p
}

//function to find the greatest common divider of the two given numbers
func gcd(x, y int) int {
   for y != 0 {
	 x, y = y, x%y 
   }
 return x 
}

//function to find the n-th Fibonacci number
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
    }
 return x
}

//Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { 
    return Fahrenheit(c*9/5 + 32) 
}

//Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { 
     return Celsius((f - 32) * 5 / 9) 
}



