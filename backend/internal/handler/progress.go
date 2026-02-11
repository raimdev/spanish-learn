package handler

import (
	"net/http"

	db "spanish-learn-backend/db/sqlc"
	"spanish-learn-backend/internal/middleware"
)

type ProgressHandler struct {
	queries *db.Queries
}

func NewProgressHandler(queries *db.Queries) *ProgressHandler {
	return &ProgressHandler{queries: queries}
}

func (h *ProgressHandler) Summary(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())

	completedLessons, _ := h.queries.CountCompletedLessons(r.Context(), userID)
	totalLessons, _ := h.queries.CountTotalLessons(r.Context())
	exercisesCompleted, _ := h.queries.CountExercisesCompleted(r.Context(), userID)
	totalExercises, _ := h.queries.CountTotalExercises(r.Context())
	dueCards, _ := h.queries.CountDueFlashcards(r.Context(), userID)

	writeJSON(w, map[string]interface{}{
		"grammar": map[string]interface{}{
			"completed": completedLessons,
			"total":     totalLessons,
		},
		"exercises": map[string]interface{}{
			"completed": exercisesCompleted,
			"total":     totalExercises,
		},
		"flashcards": map[string]interface{}{
			"due": dueCards,
		},
	})
}

func (h *ProgressHandler) ExerciseHistory(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())

	history, err := h.queries.GetExerciseHistory(r.Context(), db.GetExerciseHistoryParams{
		UserID: userID,
		Limit:  50,
	})
	if err != nil {
		writeError(w, "failed to get exercise history", http.StatusInternalServerError)
		return
	}

	type historyItem struct {
		ExerciseID  int64  `json:"exercise_id"`
		Type        string `json:"type"`
		Prompt      string `json:"prompt"`
		IsCorrect   bool   `json:"is_correct"`
		UserAnswer  string `json:"user_answer,omitempty"`
		CompletedAt string `json:"completed_at"`
	}

	items := make([]historyItem, len(history))
	for i, h := range history {
		item := historyItem{
			ExerciseID:  h.ExerciseID,
			Type:        h.Type,
			Prompt:      h.Prompt,
			IsCorrect:   h.IsCorrect == 1,
			CompletedAt: h.CompletedAt,
		}
		if h.UserAnswer.Valid {
			item.UserAnswer = h.UserAnswer.String
		}
		items[i] = item
	}
	writeJSON(w, items)
}

func (h *ProgressHandler) ReviewHistory(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())

	history, err := h.queries.GetReviewHistory(r.Context(), db.GetReviewHistoryParams{
		UserID: userID,
		Limit:  50,
	})
	if err != nil {
		writeError(w, "failed to get review history", http.StatusInternalServerError)
		return
	}

	type reviewItem struct {
		FlashcardID int64  `json:"flashcard_id"`
		Front       string `json:"front"`
		Back        string `json:"back"`
		Quality     int64  `json:"quality"`
		ReviewedAt  string `json:"reviewed_at"`
	}

	items := make([]reviewItem, len(history))
	for i, h := range history {
		items[i] = reviewItem{
			FlashcardID: h.FlashcardID,
			Front:       h.Front,
			Back:        h.Back,
			Quality:     h.Quality,
			ReviewedAt:  h.ReviewedAt,
		}
	}
	writeJSON(w, items)
}
