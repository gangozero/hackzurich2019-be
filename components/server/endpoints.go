package server

import (
	"github.com/repa40x/hackzurich2019-be/generated/models"
)

func (s *Server) StartGame() (*models.GameDescription, error) {
	id := randStringRunes(16)

	chn := s.startWorker(id)

	gameState := &Game{
		ID:           id,
		Count:        initialCount,
		CountFish:    0,
		LocationFish: []*models.Point{},
		Status:       models.GameDescriptionStatusACTIVE,
		msgChan:      chn,
	}
	err := s.setState(id, gameState)
	if err != nil {
		return nil, err
	}

	return &models.GameDescription{
		ID:     id,
		Status: gameState.Status,
	}, nil
}

func (s *Server) PauseGame(id string) (*models.GameDescription, error) {

	gameState, err := s.getState(id)
	if err != nil {
		return nil, err
	}
	gameState.Status = models.GameDescriptionStatusPAUSED

	err = s.setState(id, gameState)
	if err != nil {
		return nil, err
	}

	return &models.GameDescription{
		ID:     id,
		Status: gameState.Status,
	}, nil
}

func (s *Server) ResumeGame(id string) (*models.GameDescription, error) {

	gameState, err := s.getState(id)
	if err != nil {
		return nil, err
	}
	gameState.Status = models.GameDescriptionStatusACTIVE
	err = s.setState(id, gameState)
	if err != nil {
		return nil, err
	}

	return &models.GameDescription{
		ID:     id,
		Status: gameState.Status,
	}, nil
}

func (s *Server) GetGameDescription(id string) (*models.GameDescription, error) {

	gameState, err := s.getROState(id)
	if err != nil {
		return nil, err
	}

	return &models.GameDescription{
		ID:     id,
		Status: gameState.Status,
	}, nil
}

func (s *Server) GetGameState(id string) (*models.GameState, error) {

	gameState, err := s.getROState(id)
	if err != nil {
		return nil, err
	}

	return &models.GameState{
		Count:     int64(gameState.Count),
		CountShip: int64(gameState.CountFish),
		Ships:     gameState.LocationFish,
	}, nil
}
