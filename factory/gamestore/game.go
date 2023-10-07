package gamestore

import "fmt"

type Game interface {
	Play()
}

type ActionGame struct {
	Name string
}

func NewActionGame(name string) *ActionGame {
	return &ActionGame{Name: name}
}

func (g *ActionGame) Play() {
	fmt.Printf("playing an action game: %s\n", g.Name)
}

type AdventureGame struct {
	Name string
}

func NewAdventureGame(name string) *AdventureGame {
	return &AdventureGame{Name: name}
}

func (g *AdventureGame) Play() {
	fmt.Printf("playing an adventure game: %s\n", g.Name)
}

type RPGGame struct {
	Name string
}

func NewRPGGame(name string) *RPGGame {
	return &RPGGame{Name: name}
}

func (g *RPGGame) Play() {
	fmt.Printf("playing an RPG game: %s\n", g.Name)
}
