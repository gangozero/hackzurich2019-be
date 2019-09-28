package server

import (
	"log"

	"github.com/repa40x/hackzurich2019-be/generated/models"
)

const (
	growRate = 0.00005
)

func (s *Server) isActive(id string) bool {
	gameState, err := s.getROState(id)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	if gameState.Status != models.GameDescriptionStatusACTIVE {
		return false
	}
	return true
}

func (s *Server) grow(id string) {
	if s.isActive(id) {
		gameState, err := s.getState(id)
		if err != nil {
			log.Println(err.Error())
		}

		gameState.Count += int(float64(gameState.Count) * growRate)

		err = s.setState(id, gameState)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
