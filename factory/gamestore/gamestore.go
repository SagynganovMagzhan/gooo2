package gamestore

import (
	"fmt"
)

type GameStore struct {
	Library []Game
}

func NewGameStore() *GameStore {
	return &GameStore{}
}

func (s *GameStore) PurchaseGame(factory GameFactory, name string) {
	game := factory.CreateGame(name)
	s.Library = append(s.Library, game)
}

func (s *GameStore) ListGames() {
	fmt.Println("Available games:")
	for _, game := range s.Library {
		fmt.Printf("- %s\n", game)
	}
}
