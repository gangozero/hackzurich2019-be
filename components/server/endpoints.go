package server

import (
	"fmt"

	"github.com/repa40x/hackzurich2019-be/generated/models"
)

func (s *Server) StartGame() (*models.GameDescription, error) {
	id := randStringRunes(16)

	s.state[id] = &Game{
		ID:     id,
		Count:  initialCount,
		Status: models.GameDescriptionStatusACTIVE,
	}

	// TODO: add lifecycle

	return &models.GameDescription{
		ID:     id,
		Status: models.GameDescriptionStatusACTIVE,
	}, nil
}

func (s *Server) PauseGame(id string) (*models.GameDescription, error) {

	gameState, ok := s.state[id]
	if !ok {
		return nil, fmt.Errorf("game with ID '%s' not found", id)
	}
	gameState.Status = models.GameDescriptionStatusPAUSED
	s.state[id] = gameState

	// TODO: add lifecycle

	return &models.GameDescription{
		ID:     id,
		Status: gameState.Status,
	}, nil
}

func (s *Server) ResumeGame(id string) (*models.GameDescription, error) {

	gameState, ok := s.state[id]
	if !ok {
		return nil, fmt.Errorf("game with ID '%s' not found", id)
	}
	gameState.Status = models.GameDescriptionStatusACTIVE
	s.state[id] = gameState

	// TODO: add lifecycle

	return &models.GameDescription{
		ID:     id,
		Status: gameState.Status,
	}, nil
}

func (s *Server) GetGameDescription(id string) (*models.GameDescription, error) {

	gameState, ok := s.state[id]
	if !ok {
		return nil, fmt.Errorf("game with ID '%s' not found", id)
	}

	return &models.GameDescription{
		ID:     id,
		Status: gameState.Status,
	}, nil
}

func (s *Server) GetGameState(id string) (*models.GameState, error) {

	gameState, ok := s.state[id]
	if !ok {
		return nil, fmt.Errorf("game with ID '%s' not found", id)
	}

	return &models.GameState{
		Count: int64(gameState.Count),
	}, nil
}
