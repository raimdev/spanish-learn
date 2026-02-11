package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	db "spanish-learn-backend/db/sqlc"
	"spanish-learn-backend/internal/middleware"
)

type GrammarHandler struct {
	queries *db.Queries
}

func NewGrammarHandler(queries *db.Queries) *GrammarHandler {
	return &GrammarHandler{queries: queries}
}

type grammarListItem struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Slug      string `json:"slug"`
	Summary   string `json:"summary"`
	SortOrder int64  `json:"sort_order"`
	Completed bool   `json:"completed"`
}

func (h *GrammarHandler) List(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())

	lessons, err := h.queries.ListGrammarLessons(r.Context())
	if err != nil {
		writeError(w, "failed to list lessons", http.StatusInternalServerError)
		return
	}

	progressList, _ := h.queries.ListUserGrammarProgress(r.Context(), userID)
	progressMap := make(map[int64]bool)
	for _, p := range progressList {
		progressMap[p.LessonID] = p.Completed == 1
	}

	items := make([]grammarListItem, len(lessons))
	for i, l := range lessons {
		items[i] = grammarListItem{
			ID:        l.ID,
			Title:     l.Title,
			Slug:      l.Slug,
			Summary:   l.Summary,
			SortOrder: l.SortOrder,
			Completed: progressMap[l.ID],
		}
	}

	writeJSON(w, items)
}

type grammarDetailResponse struct {
	ID        int64            `json:"id"`
	Title     string           `json:"title"`
	Slug      string           `json:"slug"`
	Summary   string           `json:"summary"`
	Content   string           `json:"content"`
	SortOrder int64            `json:"sort_order"`
	Examples  []grammarExample `json:"examples"`
	Completed bool             `json:"completed"`
}

type grammarExample struct {
	ID          int64  `json:"id"`
	Spanish     string `json:"spanish"`
	English     string `json:"english"`
	Explanation string `json:"explanation,omitempty"`
}

func (h *GrammarHandler) GetBySlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	userID := middleware.GetUserID(r.Context())

	lesson, err := h.queries.GetGrammarLessonBySlug(r.Context(), slug)
	if err != nil {
		writeError(w, "lesson not found", http.StatusNotFound)
		return
	}

	examples, _ := h.queries.GetGrammarExamples(r.Context(), lesson.ID)

	var completed bool
	progress, err := h.queries.GetUserGrammarProgress(r.Context(), db.GetUserGrammarProgressParams{
		UserID:   userID,
		LessonID: lesson.ID,
	})
	if err == nil {
		completed = progress.Completed == 1
	}

	exList := make([]grammarExample, len(examples))
	for i, e := range examples {
		exList[i] = grammarExample{
			ID:      e.ID,
			Spanish: e.Spanish,
			English: e.English,
		}
		if e.Explanation.Valid {
			exList[i].Explanation = e.Explanation.String
		}
	}

	writeJSON(w, grammarDetailResponse{
		ID:        lesson.ID,
		Title:     lesson.Title,
		Slug:      lesson.Slug,
		Summary:   lesson.Summary,
		Content:   lesson.Content,
		SortOrder: lesson.SortOrder,
		Examples:  exList,
		Completed: completed,
	})
}

func (h *GrammarHandler) MarkComplete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	lessonID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, "invalid lesson id", http.StatusBadRequest)
		return
	}

	userID := middleware.GetUserID(r.Context())

	err = h.queries.UpsertGrammarProgress(r.Context(), db.UpsertGrammarProgressParams{
		UserID:   userID,
		LessonID: lessonID,
	})
	if err != nil {
		writeError(w, "failed to mark complete", http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]string{"message": "lesson marked as complete"})
}
