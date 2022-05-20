package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nvzard/casino-royale/service"
)

// CreateDeck Controller
func CreateDeck(c *gin.Context) {
	shuffle := c.DefaultQuery("shuffle", "false") == "true" // Shuffle query paramater
	customCardsQuery := c.DefaultQuery("cards", "")         // Custom cards query parameter

	var cards []string
	var err error
	if len(customCardsQuery) > 0 {
		cards, err = service.CreateCustomCardSequence(customCardsQuery)
	} else {
		cards = service.CreateDefaultCardSequence()
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	deck, err := service.CreateDeck(cards, shuffle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, deck.ToCreatedDeckJSON())
}

// OpenDeck Controller
func OpenDeck(c *gin.Context) {
	deckID := c.Params.ByName("deck_id")

	deck, err := service.OpenDeck(deckID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Deck not found.",
		})
		return
	}
	c.JSON(http.StatusOK, deck.ToOpenDeckJSON())
}

// Draw Controller
func Draw(c *gin.Context) {
	count, err := strconv.Atoi(c.DefaultQuery("count", "1"))
	deckID := c.Params.ByName("deck_id")

	if err != nil || count < 1 {
		c.JSON(http.StatusBadRequest, "Wrong count parameter")
		return
	}

	deck, err := service.OpenDeck(deckID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Deck not found.",
		})
		return
	}

	drawnCards := service.DrawCards(deck, count)

	c.JSON(http.StatusOK, gin.H{"cards": drawnCards})
}
