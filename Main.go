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
}
