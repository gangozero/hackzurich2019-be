package server

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/repa40x/hackzurich2019-be/generated/restapi/operations/game"
)

func (s *Server) GameStartGameHandler() game.StartGameHandler {
	return game.StartGameHandlerFunc(func(params game.StartGameParams) middleware.Responder {
		resp, err := s.StartGame()
		if err != nil {
			log.Printf("[StartGame] Error: %s", err.Error())
			return game.NewStartGameDefault(500)
		}
		return game.NewStartGameOK().WithPayload(resp)
	})
}

func (s *Server) GameGetGameDescriptionHandler() game.GetGameDescriptionHandler {
	return game.GetGameDescriptionHandlerFunc(func(params game.GetGameDescriptionParams) middleware.Responder {
		resp, err := s.GetGameDescription(params.GameID)
		if err != nil {
			log.Printf("[GetGameDescription] Error: %s", err.Error())
			return game.NewGetGameDescriptionDefault(500)
		}
		return game.NewGetGameDescriptionOK().WithPayload(resp)
	})
}

func (s *Server) GamePauseGameHandler() game.PauseGameHandler {
	return game.PauseGameHandlerFunc(func(params game.PauseGameParams) middleware.Responder {
		resp, err := s.PauseGame(params.GameID)
		if err != nil {
			log.Printf("[PauseGame] Error: %s", err.Error())
			return game.NewPauseGameDefault(500)
		}
		return game.NewPauseGameOK().WithPayload(resp)
	})
}

func (s *Server) GameResumeGameHandler() game.ResumeGameHandler {
	return game.ResumeGameHandlerFunc(func(params game.ResumeGameParams) middleware.Responder {
		resp, err := s.ResumeGame(params.GameID)
		if err != nil {
			log.Printf("[ResumeGame] Error: %s", err.Error())
			return game.NewPauseGameDefault(500)
		}
		return game.NewResumeGameOK().WithPayload(resp)
	})
}

func (s *Server) GameGetGameStateHandler() game.GetGameStateHandler {
	return game.GetGameStateHandlerFunc(func(params game.GetGameStateParams) middleware.Responder {
		resp, err := s.GetGameState(params.GameID)
		if err != nil {
			log.Printf("[GetGameState] Error: %s", err.Error())
			return game.NewGetGameStateDefault(500)
		}
		return game.NewGetGameStateOK().WithPayload(resp)
	})
}

func (s *Server) GameDestroyDisasterHandler() game.DestroyDisasterHandler {
	return game.DestroyDisasterHandlerFunc(func(params game.DestroyDisasterParams) middleware.Responder {
		resp, err := s.DestroyDisaster(params.GameID, params.Goal)
		if err != nil {
			log.Printf("[DestroyDisaster] Error: %s", err.Error())
			return game.NewDestroyDisasterDefault(500)
		}
		return game.NewDestroyDisasterOK().WithPayload(resp)
	})
}
