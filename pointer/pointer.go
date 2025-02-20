package main

import "fmt"

func main() {


	// // first method
	// myNumber:= 10;

	// // here it create pointer and also reference to the memory address of myNumber.
	// onlyNumber := &myNumber;

	// // here it dereference the memory address of myNumber and print the value of myNumber.
	// fmt.Println("myNumber: ", *onlyNumber);


	// // second method

	// var myNumber2 int = 20;
	// // here it create pointer and also reference to the memory address of myNumber2.
	// var onlyNumber2 *int = &myNumber2;

	// // here it dereference the memory address of myNumber2 and print the value of myNumber2.
	// fmt.Println("myNumber2: ", *onlyNumber2);

	i := 10;
	j := &i;

	fmt.Println("i: ", *j);


	// third method to declare pointer

	myName := "Rahul";
	var copyName = &myName;

	fmt.Println("myName: ", *copyName);


}


