package server

import (
	"fmt"
	"log"
	"time"
)

func (s *Server) startWorker(id string) chan *command {
	chn := make(chan *command)
	go func() {
		for cmd := range chn {
			err := s.doCmd(id, cmd)
			if err != nil {
				log.Printf("[Worker][id=%s] error executing command: %s", id, err.Error())
			}
		}
	}()

	// TODO: replace with better flow
	go func() {
		for {
			s.grow(id)
			s.reduceFish(id)
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			s.addFish(id)
			time.Sleep(10 * time.Second)
		}
	}()

	return chn
}

func (s *Server) doCmd(id string, cmd *command) error {
	return nil
}

func (s *Server) getState(id string) (*Game, error) {
	s.m.Lock()
	defer s.m.Unlock()
	gameState, ok := s.state[id]
	if !ok {
		return nil, fmt.Errorf("game with ID '%s' not found", id)
	}

	return gameState, nil
}

func (s *Server) getROState(id string) (*Game, error) {
	s.m.RLock()
	defer s.m.RUnlock()
	gameState, ok := s.state[id]
	if !ok {
		return nil, fmt.Errorf("game with ID '%s' not found", id)
	}
	return gameState, nil
}

func (s *Server) setState(id string, gameState *Game) error {
	s.m.Lock()
	s.state[id] = gameState
	s.m.Unlock()
	return nil
}
