package main

import (
	"fmt"
	"os"
	"time"
)

const PATH  = "./log.txt"

func main()  {
	ch := make(chan string, 100)

	createFile()


	go logging(ch)

	go func(){
		ch <- "hello"
	}()

	fmt.Scan()
}

func logging(ch chan string)  {

	for  {
		str, ok := <- ch
		if ok {
			f, _ := os.OpenFile(PATH, os.O_APPEND, os.ModeAppend)
			time := time.Now().Format(time.RFC3339)
			f.WriteString(time+"["+str+"]")
			f.Close()
		}else{
			break
		}
	}
}


func createFile(){
	f, _ :=os.Create(PATH)
	f.Close()
}
