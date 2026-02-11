package handler_test

import (
	"net/http"
	"testing"

	"spanish-learn-backend/internal/testutil"
)

func TestFlashcardDecks(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/flashcards/decks")

	var decks []struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		CardCount   int64  `json:"card_count"`
	}
	testutil.MustDecodeJSON(t, resp, &decks)

	if len(decks) != 3 {
		t.Fatalf("expected 3 flashcard decks, got %d", len(decks))
	}

	for _, d := range decks {
		if d.Name == "" {
			t.Error("deck name should not be empty")
		}
		if d.CardCount == 0 {
			t.Errorf("deck %q should have cards", d.Name)
		}
	}
}

func TestFlashcardGetDeck(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/flashcards/decks/1")

	var deck struct {
		ID        int64  `json:"id"`
		Name      string `json:"name"`
		CardCount int64  `json:"card_count"`
	}
	testutil.MustDecodeJSON(t, resp, &deck)

	if deck.ID != 1 {
		t.Errorf("expected deck ID 1, got %d", deck.ID)
	}
	if deck.Name != "Basic Greetings" {
		t.Errorf("expected name 'Basic Greetings', got %q", deck.Name)
	}
	if deck.CardCount != 12 {
		t.Errorf("expected 12 cards in deck, got %d", deck.CardCount)
	}
}

func TestFlashcardGetDeck_NotFound(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/flashcards/decks/999")
	testutil.AssertStatus(t, resp, http.StatusNotFound)
	resp.Body.Close()
}

func TestFlashcardGetDueCards(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/flashcards/due?limit=5")

	var cards []struct {
		ID     int64  `json:"id"`
		DeckID int64  `json:"deck_id"`
		Front  string `json:"front"`
		Back   string `json:"back"`
	}
	testutil.MustDecodeJSON(t, resp, &cards)

	if len(cards) != 5 {
		t.Errorf("expected 5 due cards, got %d", len(cards))
	}

	for _, c := range cards {
		if c.Front == "" || c.Back == "" {
			t.Error("card should have both front and back")
		}
	}
}

func TestFlashcardGetDueCards_ByDeck(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/flashcards/due?deck_id=1&limit=50")

	var cards []struct {
		DeckID int64 `json:"deck_id"`
	}
	testutil.MustDecodeJSON(t, resp, &cards)

	for _, c := range cards {
		if c.DeckID != 1 {
			t.Errorf("all cards should be from deck 1, got deck %d", c.DeckID)
		}
	}
}

func TestFlashcardReview(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthPost(t, server, cookie, "/api/flashcards/1/review", `{"quality":4}`)

	var result struct {
		Repetitions    int64   `json:"repetitions"`
		EasinessFactor float64 `json:"easiness_factor"`
		IntervalDays   int64   `json:"interval_days"`
		NextReviewDate string  `json:"next_review_date"`
	}
	testutil.MustDecodeJSON(t, resp, &result)

	if result.Repetitions != 1 {
		t.Errorf("expected repetitions=1, got %d", result.Repetitions)
	}
	if result.IntervalDays != 1 {
		t.Errorf("first review interval should be 1, got %d", result.IntervalDays)
	}
	if result.NextReviewDate == "" {
		t.Error("next review date should not be empty")
	}
}

func TestFlashcardReview_InvalidQuality(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthPost(t, server, cookie, "/api/flashcards/1/review", `{"quality":6}`)
	testutil.AssertStatus(t, resp, http.StatusBadRequest)
	resp.Body.Close()

	resp = testutil.AuthPost(t, server, cookie, "/api/flashcards/1/review", `{"quality":-1}`)
	testutil.AssertStatus(t, resp, http.StatusBadRequest)
	resp.Body.Close()
}

func TestFlashcardReview_ReducesDueCount(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	// Get initial due count
	resp := testutil.AuthGet(t, server, cookie, "/api/progress/summary")
	var before struct {
		Flashcards struct {
			Due int64 `json:"due"`
		} `json:"flashcards"`
	}
	testutil.MustDecodeJSON(t, resp, &before)

	// Review a card with quality 4 (correct, schedules for future)
	resp = testutil.AuthPost(t, server, cookie, "/api/flashcards/1/review", `{"quality":4}`)
	resp.Body.Close()

	// Get updated due count
	resp = testutil.AuthGet(t, server, cookie, "/api/progress/summary")
	var after struct {
		Flashcards struct {
			Due int64 `json:"due"`
		} `json:"flashcards"`
	}
	testutil.MustDecodeJSON(t, resp, &after)

	if after.Flashcards.Due >= before.Flashcards.Due {
		t.Errorf("due count should decrease after review: before=%d, after=%d",
			before.Flashcards.Due, after.Flashcards.Due)
	}
}

func TestFlashcardReview_SecondReviewUpdatesProgress(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	// First review
	resp := testutil.AuthPost(t, server, cookie, "/api/flashcards/1/review", `{"quality":4}`)
	var first struct {
		Repetitions int64 `json:"repetitions"`
	}
	testutil.MustDecodeJSON(t, resp, &first)

	// Review same card again (simulating it becoming due again)
	resp = testutil.AuthPost(t, server, cookie, "/api/flashcards/1/review", `{"quality":4}`)
	var second struct {
		Repetitions int64 `json:"repetitions"`
	}
	testutil.MustDecodeJSON(t, resp, &second)

	if second.Repetitions != first.Repetitions+1 {
		t.Errorf("repetitions should increment: first=%d, second=%d",
			first.Repetitions, second.Repetitions)
	}
}
