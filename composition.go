package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

type Player struct {
	Id, GameId     int
	Name, Location string
}

func main() {
	p := Player{}
	p.Id = 42
	p.Name = "Matt"
	p.Location = "LA"
	p.GameId = 90404
	fmt.Printf("%+v", p)
}
