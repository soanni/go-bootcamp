package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func NewUser(id int, name, location string) *User {
	u := User{Id: id, Name: name, Location: location}
	return &u
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

type Player struct {
	*User
	GameId int
}

func main() {
	p := Player{NewUser(1, "Andrey", "Samara"), 31}
	fmt.Println(p.Greetings())
}
