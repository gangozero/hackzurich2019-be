package main

import (
	"log"
	"net/http"

	"github.com/dre1080/recover"
	loads "github.com/go-openapi/loads"
	"github.com/repa40x/hackzurich2019-be/components/server"
	"github.com/repa40x/hackzurich2019-be/generated/restapi"
	"github.com/repa40x/hackzurich2019-be/generated/restapi/operations"
)

func setupGlobalMiddleware(handler http.Handler) http.Handler {
	recovery := recover.New(&recover.Options{
		Log: log.Print,
	})
	return recovery(handler)
}

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	srv := server.NewServer()

	api := operations.NewHackzurich2019BeAPI(swaggerSpec)

	api.GameStartGameHandler = srv.GameStartGameHandler()
	api.GameGetGameDescriptionHandler = srv.GameGetGameDescriptionHandler()
	api.GamePauseGameHandler = srv.GamePauseGameHandler()
	api.GameResumeGameHandler = srv.GameResumeGameHandler()
	api.GameGetGameStateHandler = srv.GameGetGameStateHandler()
	api.GameDestroyDisasterHandler = srv.GameDestroyDisasterHandler()

	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = 8080
	server.EnabledListeners = []string{"http"}

	server.SetHandler(setupGlobalMiddleware(server.GetHandler()))

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
