package gamestore

type GameFactory interface {
	CreateGame(name string) Game
}

type ActionFactory struct{}

func (f *ActionFactory) CreateGame(name string) Game {
	return NewActionGame(name)
}

type AdventureFactory struct{}

func (f *AdventureFactory) CreateGame(name string) Game {
	return NewAdventureGame(name)
}

type RPGFactory struct{}

func (f *RPGFactory) CreateGame(name string) Game {
	return NewRPGGame(name)
}
