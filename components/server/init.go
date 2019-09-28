package server

type Game struct {
	ID    string
	Count string
}

type Server struct {
	state map[string]*Game
}

func NewServer() *Server {

	return &Server{
		state: map[string]*Game{},
	}
}
