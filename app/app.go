package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// New creates a new core app
func New(host string, port int, env string) (*App, error) {

	router := gin.Default()

	// TODO: Make it configurable
	router.Use(cors.Default())

	return &App{
		Host:        host,
		Port:        port,
		Environment: env,
		Router:      router,
	}, nil
}

// App is the base struct for web api  app
type App struct {
	Host        string
	Port        int
	Environment string
	Router      *gin.Engine
}

// Start starts the app
func (app *App) Start() {

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", app.Host, app.Port), app.Router)
	if err != nil {
		log.Fatal(err)
	}
}
