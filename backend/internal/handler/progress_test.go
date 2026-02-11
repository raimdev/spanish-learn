package handler_test

import (
	"testing"

	"spanish-learn-backend/internal/testutil"
)

func TestProgressSummary_Initial(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/progress/summary")

	var summary struct {
		Grammar struct {
			Completed int64 `json:"completed"`
			Total     int64 `json:"total"`
		} `json:"grammar"`
		Exercises struct {
			Completed int64 `json:"completed"`
			Total     int64 `json:"total"`
		} `json:"exercises"`
		Flashcards struct {
			Due int64 `json:"due"`
		} `json:"flashcards"`
	}
	testutil.MustDecodeJSON(t, resp, &summary)

	if summary.Grammar.Total != 5 {
		t.Errorf("expected 5 total grammar lessons, got %d", summary.Grammar.Total)
	}
	if summary.Grammar.Completed != 0 {
		t.Errorf("expected 0 completed grammar lessons initially, got %d", summary.Grammar.Completed)
	}
	if summary.Exercises.Total != 20 {
		t.Errorf("expected 20 total exercises, got %d", summary.Exercises.Total)
	}
	if summary.Exercises.Completed != 0 {
		t.Errorf("expected 0 completed exercises initially, got %d", summary.Exercises.Completed)
	}
	if summary.Flashcards.Due != 36 {
		t.Errorf("expected 36 due flashcards initially, got %d", summary.Flashcards.Due)
	}
}

func TestProgressSummary_AfterActivity(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	// Complete a grammar lesson
	resp := testutil.AuthPost(t, server, cookie, "/api/grammar/1/complete", "{}")
	resp.Body.Close()

	// Submit an exercise
	resp = testutil.AuthPost(t, server, cookie, "/api/exercises/1/submit", `{"answer":"hablo"}`)
	resp.Body.Close()

	// Review a flashcard
	resp = testutil.AuthPost(t, server, cookie, "/api/flashcards/1/review", `{"quality":4}`)
	resp.Body.Close()

	// Check summary
	resp = testutil.AuthGet(t, server, cookie, "/api/progress/summary")
	var summary struct {
		Grammar struct {
			Completed int64 `json:"completed"`
		} `json:"grammar"`
		Exercises struct {
			Completed int64 `json:"completed"`
		} `json:"exercises"`
		Flashcards struct {
			Due int64 `json:"due"`
		} `json:"flashcards"`
	}
	testutil.MustDecodeJSON(t, resp, &summary)

	if summary.Grammar.Completed != 1 {
		t.Errorf("expected 1 completed grammar lesson, got %d", summary.Grammar.Completed)
	}
	if summary.Exercises.Completed < 1 {
		t.Errorf("expected at least 1 completed exercise, got %d", summary.Exercises.Completed)
	}
	if summary.Flashcards.Due >= 36 {
		t.Errorf("due flashcards should decrease after review, got %d", summary.Flashcards.Due)
	}
}

func TestExerciseHistory(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	// Submit a few exercises
	resp := testutil.AuthPost(t, server, cookie, "/api/exercises/1/submit", `{"answer":"hablo"}`)
	resp.Body.Close()
	resp = testutil.AuthPost(t, server, cookie, "/api/exercises/2/submit", `{"answer":"wrong"}`)
	resp.Body.Close()

	// Check exercise history
	resp = testutil.AuthGet(t, server, cookie, "/api/progress/exercises")

	var history []struct {
		ExerciseID int64  `json:"exercise_id"`
		Type       string `json:"type"`
		Prompt     string `json:"prompt"`
		IsCorrect  bool   `json:"is_correct"`
		UserAnswer string `json:"user_answer"`
	}
	testutil.MustDecodeJSON(t, resp, &history)

	if len(history) != 2 {
		t.Fatalf("expected 2 history entries, got %d", len(history))
	}

	// Verify both exercises appear in history
	exerciseIDs := map[int64]bool{}
	for _, h := range history {
		exerciseIDs[h.ExerciseID] = true
	}
	if !exerciseIDs[1] || !exerciseIDs[2] {
		t.Errorf("history should contain exercises 1 and 2, got %v", exerciseIDs)
	}
}

func TestReviewHistory(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	// Review some flashcards
	resp := testutil.AuthPost(t, server, cookie, "/api/flashcards/1/review", `{"quality":5}`)
	resp.Body.Close()
	resp = testutil.AuthPost(t, server, cookie, "/api/flashcards/2/review", `{"quality":3}`)
	resp.Body.Close()

	// Check review history
	resp = testutil.AuthGet(t, server, cookie, "/api/progress/reviews")

	var history []struct {
		FlashcardID int64  `json:"flashcard_id"`
		Front       string `json:"front"`
		Back        string `json:"back"`
		Quality     int64  `json:"quality"`
		ReviewedAt  string `json:"reviewed_at"`
	}
	testutil.MustDecodeJSON(t, resp, &history)

	if len(history) != 2 {
		t.Fatalf("expected 2 review entries, got %d", len(history))
	}

	for _, h := range history {
		if h.Front == "" || h.Back == "" {
			t.Error("review history should include card front and back")
		}
		if h.ReviewedAt == "" {
			t.Error("review history should include timestamp")
		}
	}
}

func TestExerciseHistory_Empty(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/progress/exercises")

	var history []interface{}
	testutil.MustDecodeJSON(t, resp, &history)

	if len(history) != 0 {
		t.Errorf("expected empty exercise history, got %d entries", len(history))
	}
}

func TestReviewHistory_Empty(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/progress/reviews")

	var history []interface{}
	testutil.MustDecodeJSON(t, resp, &history)

	if len(history) != 0 {
		t.Errorf("expected empty review history, got %d entries", len(history))
	}
}
