package server

import (
	"math"

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
		CountFarm:    0,
		LocationFarm: []*models.Point{},
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
		CountFarm: int64(gameState.CountFarm),
		Ships:     gameState.LocationFish,
		Farms:     gameState.LocationFarm,
	}, nil
}

func (s *Server) DestroyDisaster(id string, goal *models.Point) (*models.GameState, error) {

	gameState, err := s.getState(id)
	if err != nil {
		return nil, err
	}

	for idx, point := range gameState.LocationFish {
		if isNearby(point, goal) {
			gameState.LocationFish = remove(gameState.LocationFish, idx)
			gameState.CountFish -= 1
			break
		}
	}

	for idx, point := range gameState.LocationFarm {
		if isNearby(point, goal) {
			gameState.LocationFarm = remove(gameState.LocationFarm, idx)
			gameState.CountFarm -= 1
			break
		}
	}

	err = s.setState(id, gameState)
	if err != nil {
		return nil, err
	}

	return &models.GameState{
		Count:     int64(gameState.Count),
		CountShip: int64(gameState.CountFish),
		CountFarm: int64(gameState.CountFarm),
		Ships:     gameState.LocationFish,
		Farms:     gameState.LocationFarm,
	}, nil
}

func isNearby(p1, p2 *models.Point) bool {
	return math.Abs(p1.Lat-p2.Lat)+math.Abs(p1.Lng-p2.Lng) < 0.0000001
}

func remove(s []*models.Point, i int) []*models.Point {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
