package server

import (
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/nvzard/casino-royale/controller"
	"github.com/nvzard/casino-royale/utils"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// SetupApiServer attached routes and middleware and starts the server
func SetupApiServer() *gin.Engine {
	r := gin.New()

	// Middleware
	r.Use(gin.Recovery())
	r.Use(ginzap.Ginzap(utils.Logger(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(utils.Logger(), true))

	// Root Routes
	r.GET("/", root)
	r.GET("/health", healthcheck)

	// Deck API routes
	r.POST("/deck", controller.CreateDeck)
	r.GET("/deck/:deck_id", controller.OpenDeck)
	r.POST("/deck/:deck_id/draw", controller.Draw)

	return r
}

func root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from the Casino Royale API!",
	})
}

func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"health": "OK",
	})
}
