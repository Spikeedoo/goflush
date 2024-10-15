package core

// This is different from a full "Client entity"
// This is the minimal reference to a client stored in the main game state
type GameClient struct {
	clientId string
	isAdmin  bool
}

type GameState int

// Use const + iota to simulate "enum" behavior
const (
	StateWaiting GameState = iota
	StatePlaying
	StateFinished
)

// Main game state
type Game struct {
	id              string
	clientList      []GameClient
	state           GameState
	pot             int
	currentPlayerId string
}
