package main

import(

	"fmt"
)

func main(){
	var in = make(chan int)
	var out = make(chan bool)
	go broutine(in, out)

	for i:=0; i<10 ; i++  {
		in <- i
	}

	close(in)

	<- out
}

func broutine(in chan int, out chan bool){
	fmt.Println("in routine")
	for{
		i, ok := <- in

		fmt.Println(i, ok)

		if !ok{
			break
		}

		/*sum :=0
		for j:=0; j<i ; j++  {
			sum +=j
		}*/

	}

	out <- true
}