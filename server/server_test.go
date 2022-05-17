package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nvzard/casino-royale/database"
	"github.com/nvzard/casino-royale/model"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	router = SetupApiServer()
	_ = database.Connect()
	_ = database.Prepare()

	code := m.Run()
	os.Exit(code)
}

func testSetup(t *testing.T) {
	_ = database.DB.Migrator().DropTable(&model.Deck{})
	_ = database.DB.AutoMigrate(&model.Deck{})
}

func makeRequest(method string, url string, body io.Reader, t *testing.T) (int, string) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		logger.Errorw("Failed to make request", "errror", err.Error())
	}
	router.ServeHTTP(w, req)
	return w.Code, w.Body.String()
}

func TestHealthRoute(t *testing.T) {
	statusCode, body := makeRequest("GET", "/health", nil, t)
	assert.Equal(t, 200, statusCode)
	assert.Equal(t, `{"health":"OK"}`, body)
}

func TestRootRoute(t *testing.T) {
	statusCode, body := makeRequest("GET", "/", nil, t)
	assert.Equal(t, 200, statusCode)
	assert.Equal(t, `{"message":"Hello from the Casino Royale API!"}`, body)
}

func TestCreateDeckRoute(t *testing.T) {
	testSetup(t)

	statusCode, _ := makeRequest("POST", "/deck", nil, t)
	assert.Equal(t, http.StatusCreated, statusCode)
}

func TestOpenDeckRoute(t *testing.T) {
	testSetup(t)

	// Create a new Deck
	var newDeck model.CreatedDeckJSON
	statusCode, body := makeRequest("POST", "/deck", nil, t) // Create deck
	assert.Equal(t, http.StatusCreated, statusCode)
	_ = json.Unmarshal([]byte(body), &newDeck)

	// Check if the new deck exists
	var fetchedDeck model.OpenDeckJSON
	statusCode, body = makeRequest("GET", "/deck/"+newDeck.DeckID.String(), nil, t)
	_ = json.Unmarshal([]byte(body), &fetchedDeck)

	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, fetchedDeck.DeckID, newDeck.DeckID)
}

func TestDrawDeckRoute(t *testing.T) {
	testSetup(t)

	// Create a new Deck
	var newDeck model.CreatedDeckJSON
	statusCode, body := makeRequest("POST", "/deck?cards=AS,2S", nil, t) // Create deck
	assert.Equal(t, http.StatusCreated, statusCode)
	_ = json.Unmarshal([]byte(body), &newDeck)

	// Draw a card from the deck
	var drawnCards model.CardsJSON
	statusCode, body = makeRequest("POST", "/deck/"+newDeck.DeckID.String()+"/draw?count=1", nil, t)
	_ = json.Unmarshal([]byte(body), &drawnCards)
	assert.Equal(t, http.StatusOK, statusCode)

	// Check remaining cards in the deck
	var fetchedDeck model.OpenDeckJSON
	statusCode, body = makeRequest("GET", "/deck/"+newDeck.DeckID.String(), nil, t)
	_ = json.Unmarshal([]byte(body), &fetchedDeck)

	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, newDeck.Remaining-1, fetchedDeck.Remaining)
}
