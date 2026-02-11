package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	db "spanish-learn-backend/db/sqlc"
)

type VocabularyHandler struct {
	queries *db.Queries
}

func NewVocabularyHandler(queries *db.Queries) *VocabularyHandler {
	return &VocabularyHandler{queries: queries}
}

type vocabWord struct {
	ID              int64  `json:"id"`
	Spanish         string `json:"spanish"`
	English         string `json:"english"`
	PartOfSpeech    string `json:"part_of_speech,omitempty"`
	Category        string `json:"category,omitempty"`
	ExampleSentence string `json:"example_sentence,omitempty"`
}

func toVocabWord(w db.VocabularyWord) vocabWord {
	v := vocabWord{
		ID:      w.ID,
		Spanish: w.Spanish,
		English: w.English,
	}
	if w.PartOfSpeech.Valid {
		v.PartOfSpeech = w.PartOfSpeech.String
	}
	if w.Category.Valid {
		v.Category = w.Category.String
	}
	if w.ExampleSentence.Valid {
		v.ExampleSentence = w.ExampleSentence.String
	}
	return v
}

func (h *VocabularyHandler) List(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")

	var words []db.VocabularyWord
	var err error
	if category != "" {
		words, err = h.queries.ListVocabularyByCategory(r.Context(), sql.NullString{String: category, Valid: true})
	} else {
		words, err = h.queries.ListVocabulary(r.Context())
	}
	if err != nil {
		writeError(w, "failed to list vocabulary", http.StatusInternalServerError)
		return
	}

	result := make([]vocabWord, len(words))
	for i, word := range words {
		result[i] = toVocabWord(word)
	}
	writeJSON(w, result)
}

func (h *VocabularyHandler) Categories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.queries.ListVocabularyCategories(r.Context())
	if err != nil {
		writeError(w, "failed to list categories", http.StatusInternalServerError)
		return
	}
	// categories is []sql.NullString, extract strings
	result := make([]string, 0, len(categories))
	for _, c := range categories {
		if c.Valid {
			result = append(result, c.String)
		}
	}
	writeJSON(w, result)
}

func (h *VocabularyHandler) Quiz(w http.ResponseWriter, r *http.Request) {
	countStr := r.URL.Query().Get("count")
	count := int64(10)
	if countStr != "" {
		if n, err := strconv.ParseInt(countStr, 10, 64); err == nil && n > 0 {
			count = n
		}
	}

	category := r.URL.Query().Get("category")

	var words []db.VocabularyWord
	var err error
	if category != "" {
		words, err = h.queries.GetRandomVocabularyByCategory(r.Context(), db.GetRandomVocabularyByCategoryParams{
			Category: sql.NullString{String: category, Valid: true},
			Limit:    count,
		})
	} else {
		words, err = h.queries.GetRandomVocabulary(r.Context(), count)
	}
	if err != nil {
		writeError(w, "failed to get quiz words", http.StatusInternalServerError)
		return
	}

	result := make([]vocabWord, len(words))
	for i, word := range words {
		result[i] = toVocabWord(word)
	}
	writeJSON(w, result)
}
