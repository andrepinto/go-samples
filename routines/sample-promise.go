package main

import (
	"errors"
	"fmt"
	"time"
)

func main()  {
	product := new(Product)
	product.Id=1234
	product.Name="PS4"

	GetProductPrice(product, false).Then(
		func(obj interface{}) error{
			product.Price = obj.(PriceResponse).Price
			fmt.Println(product.Name, product.Price)
			return nil
		},
		func(err error) {
			fmt.Println(err.Error())
		})

	fmt.Scanln()
}




type Product struct {
	Id int
	Name string
	Price float32
}

type PriceResponse struct {
	Status bool
	Price float32
}

func GetProductPrice(product *Product, forceError bool) *Promise{
	result := new(Promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func(){

		if forceError{
			result.failureChannel <- errors.New("Error on get price")
		}else{
			//simulate ws call
			fmt.Println("get price for "+product.Name+".....")
			time.Sleep(time.Second * 1)
			response := PriceResponse{
				Status:true,
				Price:299.99,
			}
			result.successChannel <- response
		}
	}()

	return result
}

type Promise struct {
	successChannel chan interface{}
	failureChannel chan error
}

func(p *Promise) Then(sucess func(interface{}) error, failure func(error)) *Promise{
	result := new(Promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func(){
		select {
			case obj := <- p.successChannel:
				newErr := sucess(obj)
				if newErr == nil{
					result.successChannel <- obj
				}else{
					result.failureChannel <- newErr
				}
			case err := <- p.failureChannel:
				failure(err)
				result.failureChannel <- err
		}
	}()

	return result
}