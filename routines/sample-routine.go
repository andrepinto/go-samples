package main

import(
	"fmt"
	"time"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(2)
	for i:=0; i<10; i++{
		go routine(i)
	}
	time.Sleep(time.Second * 1)
}

func routine(i int){
	fmt.Println(i)
}
