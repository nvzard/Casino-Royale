package server

import (
	"net/http"
	"time"

	"github.com/Depado/ginprom"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/nvzard/casino-royale/controller"
	"github.com/nvzard/casino-royale/utils"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// SetupApiServer attached routes and middleware and starts the server
func SetupApiServer() *gin.Engine {
	r := gin.New()

	// Gin-prometheus
	p := ginprom.New(
		ginprom.Engine(r),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)
	r.Use(p.Instrument())

	// Middleware
	r.Use(gin.Recovery())
	r.Use(ginzap.Ginzap(utils.Logger(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(utils.Logger(), true))

	// Root Routes
	r.GET("/", root)
	r.GET("/health", healthcheck)

	// Metrics for prometheus
	r.Any("/metrics/*query", gin.WrapH(promhttp.Handler()))

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
