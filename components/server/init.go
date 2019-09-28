package server

import (
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/repa40x/hackzurich2019-be/generated/models"
)

const (
	envNameTokenOnwater = "TOKEN_ONWATER"
	envNameFlagOnwater  = "ENABLE_ONWATER"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	//check what needed ENV variables are set
	_, isSet := os.LookupEnv(envNameTokenOnwater)
	if !isSet {
		log.Fatalf("Environmental variable '%s' has to be set", envNameTokenOnwater)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type Game struct {
	ID           string
	Count        int
	CountFish    int
	CountFarm    int
	Status       string
	msgChan      chan *command
	LocationFish []*models.Point
	LocationFarm []*models.Point
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
