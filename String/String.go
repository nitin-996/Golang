package main

import (
	"fmt"
	"strings"
)



func main() {
	course:= "this is the golang practice"
	part:="practice"


	// it matches exact word if a word is capital (Practice) can other in small (practice) it will not match.
	// result:=strings.Contains(course,part)
	// fmt.Println(result)			

	result:=strings.Count(course,part)
	fmt.Println(result)

new_result:=strings.ReplaceAll(course,part,"learning demo")
fmt.Println(new_result)
}