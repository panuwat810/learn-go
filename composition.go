package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s, from %s", u.Name, u.Location)
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

	println(p.Greetings())
}
