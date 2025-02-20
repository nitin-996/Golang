package main

import ( "fmt"
		"time")



func square(i int)int{
			 time.Sleep(1000*time.Millisecond)
			return i*i
		}

func main(){

	startTime := time.Now()
	fmt.Println("Start Time: ", startTime)
	
	for i:=0; i<=50; i++{
		fmt.Println(square(i))
	}
	esclape:= time.Since(startTime)
	fmt.Println("Time taken: ", esclape)	

}