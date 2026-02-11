package router

import (
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	db "spanish-learn-backend/db/sqlc"
	"spanish-learn-backend/internal/handler"
	"spanish-learn-backend/internal/middleware"
	"spanish-learn-backend/internal/service"
)

func New(queries *db.Queries, corsOrigin string) *chi.Mux {
	r := chi.NewRouter()

	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{corsOrigin},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	authService := service.NewAuthService(queries)
	authHandler := handler.NewAuthHandler(authService)
	grammarHandler := handler.NewGrammarHandler(queries)
	vocabHandler := handler.NewVocabularyHandler(queries)
	exerciseHandler := handler.NewExerciseHandler(queries)
	flashcardHandler := handler.NewFlashcardHandler(queries)
	progressHandler := handler.NewProgressHandler(queries)

	// Public routes
	r.Post("/api/login", authHandler.Login)

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth(authService))

		r.Post("/api/logout", authHandler.Logout)
		r.Get("/api/me", authHandler.Me)

		// Grammar
		r.Get("/api/grammar", grammarHandler.List)
		r.Get("/api/grammar/{slug}", grammarHandler.GetBySlug)
		r.Post("/api/grammar/{id}/complete", grammarHandler.MarkComplete)

		// Vocabulary
		r.Get("/api/vocabulary", vocabHandler.List)
		r.Get("/api/vocabulary/categories", vocabHandler.Categories)
		r.Get("/api/vocabulary/quiz", vocabHandler.Quiz)

		// Exercises
		r.Get("/api/exercises", exerciseHandler.List)
		r.Get("/api/exercises/{id}", exerciseHandler.Get)
		r.Post("/api/exercises/{id}/submit", exerciseHandler.Submit)

		// Flashcards
		r.Get("/api/flashcards/decks", flashcardHandler.ListDecks)
		r.Get("/api/flashcards/decks/{id}", flashcardHandler.GetDeck)
		r.Get("/api/flashcards/due", flashcardHandler.GetDueCards)
		r.Post("/api/flashcards/{id}/review", flashcardHandler.Review)

		// Progress
		r.Get("/api/progress/summary", progressHandler.Summary)
		r.Get("/api/progress/exercises", progressHandler.ExerciseHistory)
		r.Get("/api/progress/reviews", progressHandler.ReviewHistory)
	})

	return r
}
