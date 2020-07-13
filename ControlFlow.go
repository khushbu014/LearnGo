package main

import (
	"fmt"
)

func main() {
	populationMap := map[string]int{
		"India": 12347859,
		"US":    1234,
		"UK":    15369,
	}

	// i is only defined in if block
	if i, ok := populationMap["India"]; ok {
		fmt.Println(i)
	}

	// in case of multiple conditions in a if statement , short cicuiting takes place

	// switch case in go

	// i := 10

	switch i := 3; i {
	case 1, 5, 6:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("0")
	}

	i := 5

	// first case will be executed (different than the break concept in other languages)
	//fallthrough executes the next case irrespective of the condition

	switch {
	case i <= 10:
		fmt.Println("executed1")
		fallthrough
	case i <= 20:
		fmt.Println("executed2")
	}

	for i, j := 0, 0; i < 10; i, j = i+3, j+2 {
		fmt.Println(i, j)
	}

	// break and continue in for loop
	for {
		if i%2 < 2 {
			break
		} else {
			continue
		}
	}

	// range for arrays and maps
	s := []int{1, 2, 3, 4}
	for k, v := range s {
		fmt.Println(k, v)
	}

	for k, v := range populationMap {
		fmt.Println(k, v)
	}

}
