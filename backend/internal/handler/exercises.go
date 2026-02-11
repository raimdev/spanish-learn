package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	db "spanish-learn-backend/db/sqlc"
	"spanish-learn-backend/internal/middleware"
)

type ExerciseHandler struct {
	queries *db.Queries
}

func NewExerciseHandler(queries *db.Queries) *ExerciseHandler {
	return &ExerciseHandler{queries: queries}
}

type exerciseListItem struct {
	ID         int64  `json:"id"`
	Type       string `json:"type"`
	Difficulty string `json:"difficulty"`
	Prompt     string `json:"prompt"`
	Hint       string `json:"hint,omitempty"`
	LessonID   *int64 `json:"lesson_id,omitempty"`
}

func (h *ExerciseHandler) List(w http.ResponseWriter, r *http.Request) {
	exType := r.URL.Query().Get("type")
	difficulty := r.URL.Query().Get("difficulty")

	var exercises []db.ListExercisesRow
	var err error

	if exType != "" && difficulty != "" {
		rows, e := h.queries.ListExercisesByTypeAndDifficulty(r.Context(), db.ListExercisesByTypeAndDifficultyParams{
			Type:       exType,
			Difficulty: difficulty,
		})
		err = e
		exercises = make([]db.ListExercisesRow, len(rows))
		for i, row := range rows {
			exercises[i] = db.ListExercisesRow{
				ID:         row.ID,
				Type:       row.Type,
				Difficulty: row.Difficulty,
				Prompt:     row.Prompt,
				Hint:       row.Hint,
				LessonID:   row.LessonID,
				CreatedAt:  row.CreatedAt,
			}
		}
	} else if exType != "" {
		rows, e := h.queries.ListExercisesByType(r.Context(), exType)
		err = e
		exercises = make([]db.ListExercisesRow, len(rows))
		for i, row := range rows {
			exercises[i] = db.ListExercisesRow{
				ID:         row.ID,
				Type:       row.Type,
				Difficulty: row.Difficulty,
				Prompt:     row.Prompt,
				Hint:       row.Hint,
				LessonID:   row.LessonID,
				CreatedAt:  row.CreatedAt,
			}
		}
	} else if difficulty != "" {
		rows, e := h.queries.ListExercisesByDifficulty(r.Context(), difficulty)
		err = e
		exercises = make([]db.ListExercisesRow, len(rows))
		for i, row := range rows {
			exercises[i] = db.ListExercisesRow{
				ID:         row.ID,
				Type:       row.Type,
				Difficulty: row.Difficulty,
				Prompt:     row.Prompt,
				Hint:       row.Hint,
				LessonID:   row.LessonID,
				CreatedAt:  row.CreatedAt,
			}
		}
	} else {
		exercises, err = h.queries.ListExercises(r.Context())
	}

	if err != nil {
		writeError(w, "failed to list exercises", http.StatusInternalServerError)
		return
	}

	items := make([]exerciseListItem, len(exercises))
	for i, e := range exercises {
		item := exerciseListItem{
			ID:         e.ID,
			Type:       e.Type,
			Difficulty: e.Difficulty,
			Prompt:     e.Prompt,
		}
		if e.Hint.Valid {
			item.Hint = e.Hint.String
		}
		if e.LessonID.Valid {
			id := e.LessonID.Int64
			item.LessonID = &id
		}
		items[i] = item
	}
	writeJSON(w, items)
}

func (h *ExerciseHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, "invalid exercise id", http.StatusBadRequest)
		return
	}

	exercise, err := h.queries.GetExercise(r.Context(), id)
	if err != nil {
		writeError(w, "exercise not found", http.StatusNotFound)
		return
	}

	resp := map[string]interface{}{
		"id":         exercise.ID,
		"type":       exercise.Type,
		"difficulty": exercise.Difficulty,
		"prompt":     exercise.Prompt,
	}
	if exercise.Hint.Valid {
		resp["hint"] = exercise.Hint.String
	}
	if exercise.Options.Valid {
		var opts []string
		json.Unmarshal([]byte(exercise.Options.String), &opts)
		resp["options"] = opts
	}
	if exercise.LessonID.Valid {
		resp["lesson_id"] = exercise.LessonID.Int64
	}

	writeJSON(w, resp)
}

type submitRequest struct {
	Answer string `json:"answer"`
}

func (h *ExerciseHandler) Submit(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, "invalid exercise id", http.StatusBadRequest)
		return
	}

	var req submitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	exercise, err := h.queries.GetExercise(r.Context(), id)
	if err != nil {
		writeError(w, "exercise not found", http.StatusNotFound)
		return
	}

	userID := middleware.GetUserID(r.Context())
	correct := strings.EqualFold(strings.TrimSpace(req.Answer), strings.TrimSpace(exercise.CorrectAnswer))

	var isCorrect int64
	if correct {
		isCorrect = 1
	}

	h.queries.RecordExerciseAttempt(r.Context(), db.RecordExerciseAttemptParams{
		UserID:     userID,
		ExerciseID: id,
		IsCorrect:  isCorrect,
		UserAnswer: sql.NullString{String: req.Answer, Valid: req.Answer != ""},
	})

	writeJSON(w, map[string]interface{}{
		"correct":        correct,
		"correct_answer": exercise.CorrectAnswer,
	})
}
