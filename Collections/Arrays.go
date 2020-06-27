package main

import (
	"fmt"
)

func main() {

	arrayOperations()

	// mostly used in place of arrays because slices don't have to have the same size over their lifetime
	sliceOperations()

}

func arrayOperations() {
	numbers := [...]int{2, 1, 4}
	fmt.Println(numbers, len(numbers))
	var numbers1 [4]int
	numbers1[3] = 6
	fmt.Println(numbers1) // it will print 0 0 0 6

	// copying array (this will create a new array pointing to other memory location)
	b := numbers
	fmt.Println(numbers, b)

	//pointing to same memory location by using &
	c := &numbers
	c[0] = 100
	fmt.Println(numbers, b, c)

	create2DArray()
}

func create2DArray() {

	var arr [3][3]int
	arr[0] = [3]int{1, 2, 3}
	arr[1] = [3]int{3, 1, 2}
	arr[2] = [3]int{2, 3, 1}
	fmt.Println(arr)

}

func sliceOperations() {
	a := []int{1, 2, 3, 4, 5}

	//always points to address ( works like python list slicing)
	b := a[1:]

	//craetes a new array with larger size
	a = append(a, 10)
	fmt.Println(a, b, len(a), len(b), cap(a), cap(b))

	//making slice using make function (type, size, capacity)
	c := make([]int, 3, 100)
	//will be initialised with 0
	fmt.Println(c)
	fmt.Printf("length: %v\n", len(c))
	fmt.Printf("length: %v\n", cap(c))

	//concating 2 slice
	a = append(a, []int{6, 7, 8}...)
	fmt.Println(a)

	stackOperationInSlice(a)
}

func stackOperationInSlice(a []int) {
	//you can use slice operation for push and pop
	stackPopFirst := a[1:]
	stackPopLast := a[:(len(a) - 1)]
	stackRemoveMiddle := append(a[:4], a[5:]...)
	//this will chnage the slice 'a" and all it's references as well, so be carefull with the use
	fmt.Println(a, stackPopFirst, stackPopLast, stackRemoveMiddle)
}
