package main

import (
	"Sagynganov_ass2/factory/gamestore"
)

func main() {
	store := gamestore.NewGameStore()

	actionFactory := &gamestore.ActionFactory{}
	adventureFactory := &gamestore.AdventureFactory{}
	rpgFactory := &gamestore.RPGFactory{}

	store.PurchaseGame(actionFactory, "Warzone")
	store.PurchaseGame(adventureFactory, "League of Legends")
	store.PurchaseGame(rpgFactory, "CS2")

	store.ListGames()
}
