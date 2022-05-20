package service

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/nvzard/casino-royale/database"
	constants "github.com/nvzard/casino-royale/helpers"
	"github.com/nvzard/casino-royale/model"
	"github.com/nvzard/casino-royale/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var logger *zap.SugaredLogger

func init() {
	logger = utils.GetLogger()
}

// CreateDeck creates creates a fresh deck using custom cards and shuffle config
func CreateDeck(cards []string, shuffle bool) (model.Deck, error) {
	deck := model.Deck{DeckID: uuid.New(), Cards: cards, IsShuffled: shuffle}

	if deck.IsShuffled {
		deck.Shuffle()
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&deck)
		if err := result.Error; err != nil {
			logger.Errorw("Failed to create deck", "error", err, "deck", deck)
			return err
		}
		return nil
	})

	if err != nil {
		logger.Errorw("Failed to create deck", "error", err)
		return deck, err
	}

	return deck, nil
}

// OpenDeck returns a new deck
func OpenDeck(deckID string) (model.Deck, error) {
	var deck model.Deck

	if err := database.DB.First(&deck, "deck_id = ?", deckID).Error; err != nil {
		return deck, err
	}

	return deck, nil
}

// DrawCards draws `count` number of card from the deck
func DrawCards(deck model.Deck, count int) model.CardsJSON {
	remaining := int(deck.Remaining())
	if remaining < count {
		count = remaining
	}

	drawnCards := deck.Cards[:count] // Draw first `count` number of cards
	deck.Cards = deck.Cards[count:]  // Remove drawn cards from the original deck

	// Update the deck
	database.DB.Save(&deck)

	return model.ToCardsJSON(drawnCards)
}

// CreateDefaultCardSequence returns [`ValueSuit`] sequence for all 52 cards
func CreateDefaultCardSequence() (codes []string) {
	for _, suit := range constants.SuitsOrder {
		for _, card := range constants.CardsOrder {
			code := string(card + suit)
			codes = append(codes, code)
		}
	}

	return codes
}

func isValidCardCode(cards []string) bool {
	for _, card := range cards {
		// todo: remove code duplication
		byteSlice := []byte(card)
		suitValue := string(byteSlice[len(byteSlice)-1])
		cardValue := string(byteSlice[:len(byteSlice)-1])

		if !strings.Contains(constants.CardsString, cardValue) || !strings.Contains(constants.SuitsString, suitValue) {
			return false
		}
	}
	return true
}

// removeDuplicates remove duplicates from the slice of custom card codes
func removeDuplicates(customCodes []string) []string {
	allKeys := make(map[string]bool)
	uniqueCodes := []string{}
	for _, item := range customCodes {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			uniqueCodes = append(uniqueCodes, item)
		}
	}
	return uniqueCodes
}

// CreateCustomCardSequence returns a custom [ValueSuit] sequence
func CreateCustomCardSequence(customCodes string) (codes []string, err error) {
	codes = strings.Split(customCodes, ",")
	if isValidCardCode(codes) {
		codes = removeDuplicates(codes)
		return codes, nil
	}
	return codes, errors.New("invalid custom card codes")
}
