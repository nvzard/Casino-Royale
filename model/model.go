package model

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/lib/pq"
	constants "github.com/nvzard/casino-royale/helpers"
)

// Deck Model
type Deck struct {
	ID         uint64    `gorm:"primaryKey"`
	DeckID     uuid.UUID `gorm:"index;type:uuid;not null"`
	IsShuffled bool
	Cards      pq.StringArray `gorm:"type:varchar(255)[]"`
}

func (deck *Deck) Remaining() int {
	return len(deck.Cards)
}

func (deck *Deck) Shuffle() {
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
	deck.IsShuffled = true
}

type CardJSON struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

type CardsJSON struct {
	Cards []CardJSON `json:"cards"`
}

type CreatedDeckJSON struct {
	DeckID     uuid.UUID `json:"deck_id"`
	IsShuffled bool      `json:"shuffled"`
	Remaining  int       `json:"remaining"`
}

type OpenDeckJSON struct {
	CreatedDeckJSON
	Cards []CardJSON `json:"cards"`
}

func getCardValue(key string) string {
	value, exists := constants.Cards[key]
	if exists {
		return value
	} else {
		return key
	}
}

func convertCardJSON(card string) CardJSON {
	byteSlice := []byte(card)
	cardValue := string(byteSlice[:len(byteSlice)-1])
	suitValue := string(byteSlice[len(byteSlice)-1])

	return CardJSON{Value: getCardValue(cardValue), Suit: constants.Suits[suitValue], Code: card}
}

func getCardsJSON(cards []string) []CardJSON {
	cardsJSON := []CardJSON{}
	for _, card := range cards {
		cardsJSON = append(cardsJSON, convertCardJSON(card))
	}

	return cardsJSON
}

func ToCardsJSON(cards []string) CardsJSON {
	return CardsJSON{Cards: getCardsJSON(cards)}
}

func (deck *Deck) ToCreatedDeckJSON() CreatedDeckJSON {
	return CreatedDeckJSON{DeckID: deck.DeckID, Remaining: deck.Remaining(), IsShuffled: deck.IsShuffled}
}

func (deck *Deck) ToOpenDeckJSON() OpenDeckJSON {
	return OpenDeckJSON{deck.ToCreatedDeckJSON(), getCardsJSON(deck.Cards)}
}
