package main

import (
	//for input and output
	"fmt"
	//for string conversion
	"strconv"
)

// variable rapping
var (
	name   string = "Sapiens"
	author string = "Yuval Noah Harari"
	pages  int    = 480
)

//I ... upper case 1st letter , will be xported globaly
var I int = 1

func main() {

	variables()

	primitives()

	constants()

}

func variables() {
	// uninitialised var have value Zer0
	var x int
	fmt.Println(x)

	// assinged var and conversion
	var intToFloat float32 = float32(x)
	fmt.Printf(" %v ,%T\n", intToFloat, intToFloat)

	var intToString string = strconv.Itoa(x)
	fmt.Printf(" %v ,%T\n", intToString, intToString)

	y := 1
	fmt.Println(y)

	// conditions
	if y == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	// declared but not used error ( all variables must be used)
	// var s string = "Sapiens"
}

func primitives() {

	// you can use int8,16,32,64
	// use uint8,16,32 for unsigned integer
	var bigNumber int64 = 123456789023456
	println(bigNumber)

	var f bool = true
	n := name == "Sapiens"
	if f {
		println(n)
	}

	//floating point numbers : float34 , 64
	fno := 3.4e15
	fmt.Println(fno)

	//complex numbers : complex64, 128 ie. float32 + float32 and float64+ float64
	var com complex64 = 2 + 5.6i
	fmt.Println(com * 10)

	//real and imag function
	fmt.Println(real(com), imag(com))

	//2 types of text type : UTF8, UTF32(rune)
	//string is immutable
	//bite slice strings

	s := "this is a string "
	r := []byte(s)
	fmt.Println(r)

	//rune = int32 type
}

// enumerated constant
const (
	//iota starts with a value 0 and increments
	_ = iota
	a = iota
	b = iota
	c

	//c will be automatically assigned value
)

func constants() {

	//constant have to be initialised at compile time
	//	ie. can't have valuee of a function

	const aConstant int = 123
	fmt.Printf("%v, %T\n", aConstant, aConstant)

	fmt.Println(a, b, c)
}
