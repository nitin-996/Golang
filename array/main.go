package main
import "fmt"

func main() {

	var fruits = [3]string{"apple", "banana", "cherry"}
	fmt.Println(fruits)

	numbers := [3]int{1, 2, 3}
	fmt.Println(numbers)

	cars := [...]string{"toyota", "ford", "chevy"}
	fmt.Println(cars)
}