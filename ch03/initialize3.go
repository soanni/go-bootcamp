package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

type Player struct {
	User
	GameId int
}

func main() {
	p := Player{
		User{Id: 42, Name: "Matt", Location: "LA"},
		90404,
	}
	// fmt.Printf("Id: %d, Name: %s, Location: %s, Game id: %d\n", p.Id, p.Name, p.Location, p.GameId)
	// p.Id = 11
	// fmt.Printf("%+v\n", p)
	fmt.Println(p.Greetings())
}
