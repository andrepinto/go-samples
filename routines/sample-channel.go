package main

import(

	"fmt"
)

func main()  {

	var c = make(chan int)

	for i:=0; i<10 ; i++  {
		go myroutine(i, c)
	}

	for i:=0; i<10 ; i++  {
		fmt.Println(i, <-c)
	}
}

func myroutine(i int, c chan int){
	c <- i
}