package main

import(
	"fmt"
)

const(
	ADD_MONEY = "add_money"
	REMOVE_MONEY = "remove_money"
)

func main()  {
	account := CreateAccount()

	handlerAddMoney := make(chan string)
	handlerRemoveMoney := make(chan string)

	account.AddEventListener(ADD_MONEY, handlerAddMoney)
	account.AddEventListener(REMOVE_MONEY, handlerRemoveMoney)

	go func(){
		for {
			msg := <- handlerAddMoney
			fmt.Println(msg)
		}
	}()

	go func(){
		for {
			msg := <- handlerRemoveMoney
			fmt.Println(msg)
		}
	}()



	account.DispatchEvent(ADD_MONEY, "add money 1")
	account.DispatchEvent(REMOVE_MONEY, "remove money 1")
	account.DispatchEvent(ADD_MONEY, "add money 2")

	account.RemoveEventListener(ADD_MONEY, handlerAddMoney)
	account.RemoveEventListener(REMOVE_MONEY, handlerAddMoney)

	fmt.Scanln()
}

type Account struct {
	eventListeners map[string][]chan string
}

func CreateAccount() *Account  {
	acc := new(Account)
	acc.eventListeners = make(map[string][]chan string)
	return acc
}

func (a *Account) AddEventListener(event string, channel chan string){
	if _, present := a.eventListeners[event]; present{
		a.eventListeners[event] = append(a.eventListeners[event], channel)
	}else{
		a.eventListeners[event] = []chan string{channel}
	}
}

func (a *Account) RemoveEventListener(event string, listenerChannel chan string){
	if _, present := a.eventListeners[event]; present{
		for j, _ := range a.eventListeners[event]{
			if a.eventListeners[event][j] == listenerChannel{
				a.eventListeners[event] =
						append(a.eventListeners[event][:j], a.eventListeners[event][j+1:]...)
				break
			}
		}
	}
}

func (a *Account) DispatchEvent(event string, response string){
	if _, present := a.eventListeners[event]; present{
		for _, handler := range a.eventListeners[event]{
			go func(handler chan string){
				handler <- response
			}(handler)
		}
	}
}