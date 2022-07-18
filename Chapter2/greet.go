//This file consists of examples from the first week's class
package main

import (
	"fmt"
	"os"
)


func main() {
    fmt.Println("Hello World")
	greet("Berke") //calling greet function
	
	//calling  createGreet function
	var name1 string = "Berke"
    var greeting = createGreet(name1)
	fmt.Printf("%s", greeting)
	
	//calling  createGreet2 function
	name2 := os.Args[1]
	greeting1 := createGreet2(name2)
	fmt.Printf("%s\n", greeting1)

    //taking input from user by using Scanln Funcion
	var name3 string
    fmt.Println("Please enter your name:")
    fmt.Scanln(&name3)
    greeting2 := createGreet2(name3)
	fmt.Printf("%s\n", greeting2)
	
	//Created an instance of Person.greet() method is called to greet the person
	greeter := Person{"Mehmet"}
	var greeting3 = greeter.greet()
	fmt.Printf("%s\n", greeting3)

    //Different greeting examples 
	greetPrinter(createGreetInTurkish, "Berk")
    greetPrinter(createGreetInEnglish, "Mana")
   
    greetCreator := createGreetInTurkish 
    greetPrinter(greetCreator, "Selen")

    func(name string) {
	    greeting := "Selam " + name + " :)" 
	    fmt.Printf("%s\n", greeting)
    }("Kemal")
    closure := func(name string) { 
	    greeting := "Selam " + name + " :)" 
	    fmt.Printf("%s\n", greeting)
    }
    closure("Emirhan")
    anotherGreetPrinter(closure, "Batu")

}

//functions

func greet(name string) {
	fmt.Printf("Selam %s :)\n", name)
}

func createGreet(name1 string) string {
  greeting := "Selamlar " + name1 + " :) \n"
  return greeting
}

func createGreet2(name2 string) string {
  return  "Merhaba " +  name2 + " :)"
}


type Person struct {
	name4 string
}

//Function greet() is associated with struct Person and so it becomes a method.
func (p Person) greet() string {
	return "Selam " +  p.name4 + " :)"
}

//Different greeting functions
func createGreetInTurkish(name string) string { 
	return "Selam " + name + " :)"
}
func createGreetInEnglish(name string) string { 
	return "Hi " + name + " :)"
}
func greetPrinter(function func(it string) string, name string){ 
	var greeting = function(name)
    fmt.Printf("%s\n", greeting)
}
func anotherGreetPrinter(function func(it string), name string){ 
	function(name)
}