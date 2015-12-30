package main

import(
	"fmt"
	"time"
	"sync"
	"runtime"
)

var m sync.Mutex

func main(){
	runtime.GOMAXPROCS(2)
	accounts := map[string]float32{
		"Andre": 1000.0,
		"Pinto": 2000.0,
	}

	go add(accounts, "Andre", 4000)
	go total(accounts)

	time.Sleep(time.Second*2)
}

func add(accounts map[string]float32, name string, balance float32){
	m.Lock()
	time.Sleep(time.Second*1)
	accounts[name] = balance
	fmt.Println(accounts)
	m.Unlock()
}

func total(accounts map[string]float32){
	m.Lock()
	var tt float32

	for _, balance := range accounts{
		tt += balance
	}

	fmt.Println(tt)
	m.Unlock()
}