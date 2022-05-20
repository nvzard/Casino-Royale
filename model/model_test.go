package model

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetRemaining(t *testing.T) {
	assert := assert.New(t)

	cards := []string{"AC", "2S", "4C"}
	deck := Deck{DeckID: uuid.New(), IsShuffled: false, Cards: cards}
	assert.Equal(deck.Remaining(), 3)
}

func TestShuffle(t *testing.T) {
	assert := assert.New(t)

	cards := []string{"2S", "4C"}
	deck := Deck{DeckID: uuid.New(), IsShuffled: false, Cards: cards}
	deck.Shuffle()
	assert.NotEqual(deck.Cards, cards)
}

func TestToCardsJSON(t *testing.T) {
	assert := assert.New(t)

	deck := Deck{DeckID: uuid.New(), IsShuffled: false, Cards: []string{"2D", "6C"}}
	cardsJSON := CardsJSON{
		Cards: []CardJSON{
			{
				Value: "2",
				Suit:  "DIAMONDS",
				Code:  "2D",
			},
			{
				Value: "6",
				Suit:  "CLUBS",
				Code:  "6C",
			},
		},
	}
	assert.Equal(ToCardsJSON(deck.Cards), cardsJSON)
}

func TestToCreatedDeckJSON(t *testing.T) {
	assert := assert.New(t)
	uuid := uuid.New()
	deck := Deck{DeckID: uuid, IsShuffled: false, Cards: []string{"2S"}}
	deckJSON := CreatedDeckJSON{
		DeckID:     uuid,
		IsShuffled: false,
		Remaining:  1,
	}
	assert.Equal(deck.ToCreatedDeckJSON(), deckJSON)
}

func TestToOpenDeckJSON(t *testing.T) {
	assert := assert.New(t)

	uuid := uuid.New()
	deck := Deck{DeckID: uuid, IsShuffled: false, Cards: []string{"2S"}}
	deckJSON := OpenDeckJSON{
		CreatedDeckJSON: CreatedDeckJSON{
			DeckID:     uuid,
			IsShuffled: false,
			Remaining:  1,
		},
		Cards: []CardJSON{
			{
				Value: "2",
				Suit:  "SPADES",
				Code:  "2S",
			},
		},
	}
	assert.Equal(deck.ToOpenDeckJSON(), deckJSON)
}
