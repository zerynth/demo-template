// @title IngestionService API
// @version 1.0
// @description IngManager Service manage the data ingestion from ZDM
// @termsOfService http://swagger.io/terms/

// @contact.name Zerynth
// @contact.url http://www.zerynth.com
// @contact.email zerynth@zerynth.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080

package main

import (
	"net/http"
	"os"

	"ingestion/database"
	"ingestion/service"

	"github.com/gorilla/handlers"

	"github.com/go-kit/kit/log"
	_ "github.com/lib/pq"
)

type serverConfigurations struct {
	Port string `default:":80"`
}

var serverConfig serverConfigurations
var tsConfig database.TSConfig

func init() {
	//env configuration
	service.LoadIntoWithPrefix(&serverConfig, "SERVER")
	service.LoadIntoWithPrefix(&tsConfig, "TS")
}

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))

	logger.Log("DB host", tsConfig.Host, "port", tsConfig.Port, "db name", tsConfig.Name, "db user", tsConfig.User, "pass", tsConfig.Password)
	logger.Log("Server config: port", serverConfig.Port)

	db, err := database.ConnectToDB(tsConfig)
	if err != nil {
		logger.Log("err", "Can't connect to database", "err", err.Error())
		return
	}

	repository := database.NewModelRepository(db)
	ingManagerService := service.NewIngManagerService(repository)
	ingManagerService = service.LoggingMiddleware(logger)(ingManagerService)
	ingManagerHandler := service.MakeHTTPHandler(ingManagerService, logger)

	ingManagerHandler = handlers.LoggingHandler(os.Stdout, ingManagerHandler)

	logger.Log("Service", "started", "port", serverConfig.Port)
	logger.Log(http.ListenAndServe(serverConfig.Port, ingManagerHandler))
}
