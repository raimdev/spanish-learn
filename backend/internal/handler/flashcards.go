package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	db "spanish-learn-backend/db/sqlc"
	"spanish-learn-backend/internal/middleware"
	"spanish-learn-backend/internal/service"
)

type FlashcardHandler struct {
	queries *db.Queries
}

func NewFlashcardHandler(queries *db.Queries) *FlashcardHandler {
	return &FlashcardHandler{queries: queries}
}

type deckResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	CardCount   int64  `json:"card_count"`
}

func (h *FlashcardHandler) ListDecks(w http.ResponseWriter, r *http.Request) {
	decks, err := h.queries.ListFlashcardDecks(r.Context())
	if err != nil {
		writeError(w, "failed to list decks", http.StatusInternalServerError)
		return
	}

	result := make([]deckResponse, len(decks))
	for i, d := range decks {
		count, _ := h.queries.CountCardsInDeck(r.Context(), d.ID)
		result[i] = deckResponse{
			ID:        d.ID,
			Name:      d.Name,
			CardCount: count,
		}
		if d.Description.Valid {
			result[i].Description = d.Description.String
		}
	}
	writeJSON(w, result)
}

func (h *FlashcardHandler) GetDeck(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, "invalid deck id", http.StatusBadRequest)
		return
	}

	deck, err := h.queries.GetFlashcardDeck(r.Context(), id)
	if err != nil {
		writeError(w, "deck not found", http.StatusNotFound)
		return
	}

	count, _ := h.queries.CountCardsInDeck(r.Context(), id)

	resp := deckResponse{
		ID:        deck.ID,
		Name:      deck.Name,
		CardCount: count,
	}
	if deck.Description.Valid {
		resp.Description = deck.Description.String
	}
	writeJSON(w, resp)
}

type flashcardResponse struct {
	ID     int64  `json:"id"`
	DeckID int64  `json:"deck_id"`
	Front  string `json:"front"`
	Back   string `json:"back"`
}

func (h *FlashcardHandler) GetDueCards(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())

	limitStr := r.URL.Query().Get("limit")
	limit := int64(20)
	if limitStr != "" {
		if n, err := strconv.ParseInt(limitStr, 10, 64); err == nil && n > 0 {
			limit = n
		}
	}

	deckIDStr := r.URL.Query().Get("deck_id")

	var cards []db.Flashcard
	var err error

	if deckIDStr != "" {
		deckID, parseErr := strconv.ParseInt(deckIDStr, 10, 64)
		if parseErr != nil {
			writeError(w, "invalid deck_id", http.StatusBadRequest)
			return
		}
		cards, err = h.queries.GetDueFlashcardsByDeck(r.Context(), db.GetDueFlashcardsByDeckParams{
			UserID: userID,
			DeckID: deckID,
			Limit:  limit,
		})
	} else {
		cards, err = h.queries.GetDueFlashcards(r.Context(), db.GetDueFlashcardsParams{
			UserID: userID,
			Limit:  limit,
		})
	}

	if err != nil {
		writeError(w, "failed to get due cards", http.StatusInternalServerError)
		return
	}

	result := make([]flashcardResponse, len(cards))
	for i, c := range cards {
		result[i] = flashcardResponse{
			ID:     c.ID,
			DeckID: c.DeckID,
			Front:  c.Front,
			Back:   c.Back,
		}
	}
	writeJSON(w, result)
}

type reviewRequest struct {
	Quality int `json:"quality"`
}

func (h *FlashcardHandler) Review(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	cardID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, "invalid flashcard id", http.StatusBadRequest)
		return
	}

	var req reviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.Quality < 0 || req.Quality > 5 {
		writeError(w, "quality must be between 0 and 5", http.StatusBadRequest)
		return
	}

	userID := middleware.GetUserID(r.Context())

	// Get current progress or use defaults
	var reps int64
	var ef float64 = 2.5
	var interval int64

	progress, err := h.queries.GetFlashcardProgress(r.Context(), db.GetFlashcardProgressParams{
		UserID:      userID,
		FlashcardID: cardID,
	})
	if err == nil {
		reps = progress.Repetitions
		ef = progress.EasinessFactor
		interval = progress.IntervalDays
	}

	result := service.SM2(req.Quality, reps, ef, interval)

	h.queries.UpsertFlashcardProgress(r.Context(), db.UpsertFlashcardProgressParams{
		UserID:         userID,
		FlashcardID:    cardID,
		Repetitions:    result.Repetitions,
		EasinessFactor: result.EasinessFactor,
		IntervalDays:   result.IntervalDays,
		NextReviewDate: result.NextReviewDate,
	})

	h.queries.CreateReviewLog(r.Context(), db.CreateReviewLogParams{
		UserID:      userID,
		FlashcardID: cardID,
		Quality:     int64(req.Quality),
	})

	writeJSON(w, map[string]interface{}{
		"repetitions":      result.Repetitions,
		"easiness_factor":  result.EasinessFactor,
		"interval_days":    result.IntervalDays,
		"next_review_date": result.NextReviewDate,
	})
}
