package main

import "fmt"

//Books ...
type Books struct {
	name    string
	price   float32
	reviews []string
}

func main() {

	createMap()

	createStruct()

}

func createMap() {

	mapStateNumber := make(map[string]int)
	mapStateNumber = map[string]int{
		"U.P.":        12,
		"M.P.":        17,
		"Maharashtra": 18,
	}

	//adding a key
	//map is unordered
	mapStateNumber["Karnataka"] = 21

	//delete a key
	//after deletion the value of key is assigned 0
	delete(mapStateNumber, "U.P.")

	//to check if it exists use following
	_, ok := mapStateNumber["U.P."]
	// ok will return false if it does not exists in map
	fmt.Println(ok)

	fmt.Println(mapStateNumber, len(mapStateNumber))
	fmt.Println(mapStateNumber["M.P."])

	// map is used and sent through reference when passed in a function

}

func createStruct() {
	aBook := Books{
		name:  "To kill a mocking bird",
		price: 400.10,
		reviews: []string{
			"Good",
			"Very well wirtten",
			"Must read",
		},
	}

	fmt.Println(aBook)
}
