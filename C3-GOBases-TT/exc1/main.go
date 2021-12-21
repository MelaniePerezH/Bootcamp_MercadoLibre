package main

import "fmt"

type User struct {
	Name     string
	Lastname string
	Email    string
	Password string
	Age      int
}

func setName(user *User, newName string, newLastname string) {
	user.Name = newName
	user.Lastname = newLastname
}

func setEmail(user *User, newEmail string) {
	user.Email = newEmail
}

func main() {
	test_ := User{
		Name:     "Melanie",
		Email:    "test@test.org",
		Password: "password",
		Age:      21,
	}
	fmt.Println(test_)
	setName(&test_, "Juliett", "Perez")
	fmt.Println(test_)
	setEmail(&test_, "test_2@test.org.co")
	fmt.Println(test_)
}
