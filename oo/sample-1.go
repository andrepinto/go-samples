package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func (p Person) GetPersonInfo() string{
	return fmt.Sprintf("Name: %s, Age: %d ", p.Name, p.Age)
}

type User struct {
	Person
	Username string
	Password string
}

func (u User) GetUserInfo() string  {
	return fmt.Sprintf("User: %s, Username: %s", u.Person.Name, u.Username)
}

func (u *User) ChangeUserPassword(password string) {
	u.Password = password
}

func main(){

	fmt.Println("----------------------------")
	fmt.Println("CREATE PERSON")
	fmt.Println("----------------------------")
	person := Person{
		Name:"Andre",
		Age: 29,
	}

	fmt.Println(person.GetPersonInfo())

	fmt.Println("----------------------------")
	fmt.Println("CREATE USER EXTEND PERSON")
	fmt.Println("----------------------------")

	var user User = User{}
	user.Person = person
	user.Username = "andre.pinto"
	user.Password = "123456"

	fmt.Println(user.GetPersonInfo())
	fmt.Println(user.GetUserInfo())

	fmt.Println("----------------------------")
	fmt.Println("CHANGE PASSWORD")
	fmt.Println("----------------------------")

	fmt.Println("Old Password:", user.Password)
	user.ChangeUserPassword("45678")
	fmt.Println("New Password:", user.Password)

}
