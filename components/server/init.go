package server

import (
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const initialCount = 1000000

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//TODO: add lock for count
type Game struct {
	ID      string
	Count   int
	Status  string
	msgChan chan *command
}

type Server struct {
	state map[string]*Game
	m     sync.RWMutex
}

func NewServer() *Server {

	return &Server{
		state: map[string]*Game{},
	}
}

type command struct {
	action string
}
