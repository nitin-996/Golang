package main

import (
	"fmt" 
	"time"
)

func sell(ch chan string){
	ch <- "Hello"
fmt.Println("Value sent")
}

func main(){
	ch := make(chan string)
	go sell(ch)
	go buy(ch)
	time.Sleep(1000*time.Millisecond)
}	

func buy(ch chan string){

	fmt.Print("waiting for the value")
	val :=<- ch 

	fmt.Println(val)
}		
