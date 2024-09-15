package basics

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// Variables
func variables() {
	var population int16 = 32767 // 2^15 - 1
	population = population + 1  // overflow
	fmt.Println(population)

	// It's a good practice to think about the variable type to optimize memory usage
	var red uint8 = 255
	var green uint8 = 123
	var blue uint8 = 111

	var name string = "Cristovão\nNeto"
	var name2 string = `Cristovao
	Neto`

	// The len function returns the number of bytes. Strings are encoded using UTF-8
	fmt.Println(len(name))
	fmt.Println(len("à"))
	fmt.Println(len("a"))

	// Get string length
	fmt.Println(utf8.RuneCountInString("à"))

	// Creates a variable with a default value
	var defaultValue int
	var omitType = "string"
	dropVar := "string"

	a, b, c := 1, 2, 3

	// It's a good practice to declare types whenever it's not obvious
	// notObviousType := foo()

	// You can use const instead of var
	const pi float32 = 3.1415

	// Avoid compilation errors
	fmt.Println(red + green + blue)
	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(omitType)
	fmt.Println(dropVar)
	fmt.Println(a + b + c)
	fmt.Println(defaultValue)

}

// Functions

func division(a int, b int) (int, int, error) {
	if b == 0 {
		error := errors.New("Cannot divide by 0")
		return 0, 0, error
	}
	var result int = a / b
	var remainder int = a % b
	return result, remainder, nil
}

// control structures

func switchStatement() {
	var age uint8 = 21
	var canDrive bool = age > 18

	// Conditional switch
	switch canDrive {
	case true:
		fmt.Println("You can drive")
	case false:
		fmt.Println("You cannot drive")
	}

	// Default switch
	switch {
	case canDrive:
		fmt.Println("You can drive")
	case !canDrive:
		fmt.Println("You cannot drive")
	}
}

// arrays

func arrays() {
	var intArray [3]int32 = [3]int32{1, 2, 3}
	// var intArray [3]int32 = [...]int32{1, 2, 3}
	// intArray = [...]int32{1, 2, 3}
	intArray[1] = 123

	fmt.Println(intArray[0], intArray[1:3])
	// Print memory address
	fmt.Println(&intArray[0], &intArray[1], &intArray[2])

}

// slices
// the length of a slice can grow and shrink as you see fit.

func slices() {
	var intSlice []int32 = []int32{1, 2, 3}
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 8)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
}

// maps

func maps() {
	// var myMap map[string]uint8 = make(map[string]uint8)
	var ages = map[string]uint8{"Cristovao": 21, "Savio": 28}
	var age, ok = ages["InvalidKey"]
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("Invalid key")
	}
	delete(ages, "Cristovao")
	fmt.Println(ages)

}

// loops

func loops() {
	var ages = map[string]uint8{"Cristovao": 21, "Savio": 28}
	for name, age := range ages {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for index, value := range [3]int32{12, 13, 33} {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
