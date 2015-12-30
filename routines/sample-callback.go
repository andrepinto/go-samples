package main

import (
	"fmt"
	"time"
)


func main()  {
	itm := CreateItem()
	fmt.Println(itm.Name, itm.Price)
}

func CreateItem() Item{
	item := new(Item)
	item.Id=1234
	item.Name="PS4"

	ch := make(chan float32)

	go GetPrice(item, ch)

	item.Price = <-ch

	return *item

}

type Item struct {
	Id int
	Name string
	Price float32
}

func GetPrice(item *Item, cb chan float32){

	//simulate ws call
	fmt.Println("get price for "+item.Name+".....")
	time.Sleep(time.Second * 1)
	cb <- 299.99
}
