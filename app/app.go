package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// New creates a new core app
func New(config *viper.Viper) (*App, error) {

	router := gin.Default()

	// TODO: Make it configurable
	router.Use(cors.Default())

	return &App{
		Router: router,
		Config: config,
	}, nil
}

// App is the base struct for web api  app
type App struct {
	Router *gin.Engine
	Config *viper.Viper
}

// Start starts the app
func (app *App) Start() {
	host := app.Config.GetString("http.host")
	port := app.Config.GetInt("http.port")

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), app.Router)
	if err != nil {
		log.Fatal(err)
	}
}
