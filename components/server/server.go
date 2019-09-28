package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/repa40x/hackzurich2019-be/generated/restapi/operations/game"
)

func (s *Server) GameStartGameHandler() game.StartGameHandler {
	return game.StartGameHandlerFunc(func(params game.StartGameParams) middleware.Responder {
		return middleware.NotImplemented("operation game.StartGame has to be implemented")
	})
}

func (s *Server) GameGetGameDescriptionHandler() game.GetGameDescriptionHandler {
	return game.GetGameDescriptionHandlerFunc(func(params game.GetGameDescriptionParams) middleware.Responder {
		return middleware.NotImplemented("operation game.GetGameDescription has to be implemented")
	})
}

func (s *Server) GamePauseGameHandler() game.PauseGameHandler {
	return game.PauseGameHandlerFunc(func(params game.PauseGameParams) middleware.Responder {
		return middleware.NotImplemented("operation game.PauseGame has to be implemented")
	})
}

func (s *Server) GameResumeGameHandler() game.ResumeGameHandler {
	return game.ResumeGameHandlerFunc(func(params game.ResumeGameParams) middleware.Responder {
		return middleware.NotImplemented("operation game.ResumeGame has to be implemented")
	})
}

func (s *Server) GameGetGameStateHandler() game.GetGameStateHandler {
	return game.GetGameStateHandlerFunc(func(params game.GetGameStateParams) middleware.Responder {
		return middleware.NotImplemented("operation game.GetGameState has to be implemented")
	})
}

func (s *Server) GameDestroyDisasterHandler() game.DestroyDisasterHandler {
	return game.DestroyDisasterHandlerFunc(func(params game.DestroyDisasterParams) middleware.Responder {
		return middleware.NotImplemented("operation game.DestroyDisaster has to be implemented")
	})
}
