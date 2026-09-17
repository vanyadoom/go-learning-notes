package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Active bool
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

func (u *User) Activate() {
	u.Active = true
}

func (u *User) Deactivate() {
	u.Active = false
	return
}

func (u User) Status() string {
	if u.Active == true {
		return "Активен"
	}
	return "Неактивен"
}

func main() {
	var name string
	var age int
	fmt.Scan(&name, &age)
	user := User{
		Name:   name,
		Age:    age,
		Active: false,
	}
	fmt.Println("Совершеннолетний: ", user.IsAdult())
	fmt.Println("начальный статус: ", user.Status())

	user.Activate()
	fmt.Println("После активации: ", user.Status())

	user.Deactivate()
	fmt.Println("После деактивации: ", user.Status())
}
