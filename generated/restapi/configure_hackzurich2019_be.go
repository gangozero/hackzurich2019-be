// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/repa40x/hackzurich2019-be/generated/restapi/operations"
	"github.com/repa40x/hackzurich2019-be/generated/restapi/operations/game"
)

//go:generate swagger generate server --target ../../generated --name Hackzurich2019Be --spec ../../openapi/swagger.yaml --exclude-main

func configureFlags(api *operations.Hackzurich2019BeAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.Hackzurich2019BeAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.GameDestroyDisasterHandler == nil {
		api.GameDestroyDisasterHandler = game.DestroyDisasterHandlerFunc(func(params game.DestroyDisasterParams) middleware.Responder {
			return middleware.NotImplemented("operation game.DestroyDisaster has not yet been implemented")
		})
	}
	if api.GameGetGameDescriptionHandler == nil {
		api.GameGetGameDescriptionHandler = game.GetGameDescriptionHandlerFunc(func(params game.GetGameDescriptionParams) middleware.Responder {
			return middleware.NotImplemented("operation game.GetGameDescription has not yet been implemented")
		})
	}
	if api.GameGetGameStateHandler == nil {
		api.GameGetGameStateHandler = game.GetGameStateHandlerFunc(func(params game.GetGameStateParams) middleware.Responder {
			return middleware.NotImplemented("operation game.GetGameState has not yet been implemented")
		})
	}
	if api.GamePauseGameHandler == nil {
		api.GamePauseGameHandler = game.PauseGameHandlerFunc(func(params game.PauseGameParams) middleware.Responder {
			return middleware.NotImplemented("operation game.PauseGame has not yet been implemented")
		})
	}
	if api.GameResumeGameHandler == nil {
		api.GameResumeGameHandler = game.ResumeGameHandlerFunc(func(params game.ResumeGameParams) middleware.Responder {
			return middleware.NotImplemented("operation game.ResumeGame has not yet been implemented")
		})
	}
	if api.GameStartGameHandler == nil {
		api.GameStartGameHandler = game.StartGameHandlerFunc(func(params game.StartGameParams) middleware.Responder {
			return middleware.NotImplemented("operation game.StartGame has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
