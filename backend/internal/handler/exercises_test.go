package handler_test

import (
	"net/http"
	"testing"

	"spanish-learn-backend/internal/testutil"
)

func TestExercisesList(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/exercises")

	var exercises []struct {
		ID         int64  `json:"id"`
		Type       string `json:"type"`
		Difficulty string `json:"difficulty"`
		Prompt     string `json:"prompt"`
	}
	testutil.MustDecodeJSON(t, resp, &exercises)

	if len(exercises) != 20 {
		t.Fatalf("expected 20 exercises, got %d", len(exercises))
	}
}

func TestExercisesList_FilterByType(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/exercises?type=fill_blank")

	var exercises []struct {
		Type string `json:"type"`
	}
	testutil.MustDecodeJSON(t, resp, &exercises)

	if len(exercises) == 0 {
		t.Fatal("should have fill_blank exercises")
	}
	for _, ex := range exercises {
		if ex.Type != "fill_blank" {
			t.Errorf("expected type 'fill_blank', got %q", ex.Type)
		}
	}
}

func TestExercisesList_FilterByDifficulty(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/exercises?difficulty=beginner")

	var exercises []struct {
		Difficulty string `json:"difficulty"`
	}
	testutil.MustDecodeJSON(t, resp, &exercises)

	if len(exercises) == 0 {
		t.Fatal("should have beginner exercises")
	}
	for _, ex := range exercises {
		if ex.Difficulty != "beginner" {
			t.Errorf("expected difficulty 'beginner', got %q", ex.Difficulty)
		}
	}
}

func TestExercisesList_FilterByTypeAndDifficulty(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/exercises?type=conjugation&difficulty=intermediate")

	var exercises []struct {
		Type       string `json:"type"`
		Difficulty string `json:"difficulty"`
	}
	testutil.MustDecodeJSON(t, resp, &exercises)

	for _, ex := range exercises {
		if ex.Type != "conjugation" || ex.Difficulty != "intermediate" {
			t.Errorf("expected conjugation/intermediate, got %s/%s", ex.Type, ex.Difficulty)
		}
	}
}

func TestExerciseGet(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/exercises/1")

	var exercise struct {
		ID         int64  `json:"id"`
		Type       string `json:"type"`
		Difficulty string `json:"difficulty"`
		Prompt     string `json:"prompt"`
		Hint       string `json:"hint"`
	}
	testutil.MustDecodeJSON(t, resp, &exercise)

	if exercise.ID != 1 {
		t.Errorf("expected exercise ID 1, got %d", exercise.ID)
	}
	if exercise.Prompt == "" {
		t.Error("exercise prompt should not be empty")
	}
}

func TestExerciseGet_NotFound(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/exercises/999")
	testutil.AssertStatus(t, resp, http.StatusNotFound)
	resp.Body.Close()
}

func TestExerciseSubmit_CorrectAnswer(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	// Exercise 1: "Yo ___ (hablar) espanol." -> "hablo"
	resp := testutil.AuthPost(t, server, cookie, "/api/exercises/1/submit", `{"answer":"hablo"}`)

	var result struct {
		Correct       bool   `json:"correct"`
		CorrectAnswer string `json:"correct_answer"`
	}
	testutil.MustDecodeJSON(t, resp, &result)

	if !result.Correct {
		t.Error("answer 'hablo' should be correct for exercise 1")
	}
	if result.CorrectAnswer != "hablo" {
		t.Errorf("expected correct answer 'hablo', got %q", result.CorrectAnswer)
	}
}

func TestExerciseSubmit_IncorrectAnswer(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthPost(t, server, cookie, "/api/exercises/1/submit", `{"answer":"hablas"}`)

	var result struct {
		Correct       bool   `json:"correct"`
		CorrectAnswer string `json:"correct_answer"`
	}
	testutil.MustDecodeJSON(t, resp, &result)

	if result.Correct {
		t.Error("answer 'hablas' should be incorrect for exercise 1")
	}
	if result.CorrectAnswer != "hablo" {
		t.Errorf("should return correct answer, got %q", result.CorrectAnswer)
	}
}

func TestExerciseSubmit_CaseInsensitive(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthPost(t, server, cookie, "/api/exercises/1/submit", `{"answer":"Hablo"}`)

	var result struct {
		Correct bool `json:"correct"`
	}
	testutil.MustDecodeJSON(t, resp, &result)

	if !result.Correct {
		t.Error("answer checking should be case-insensitive")
	}
}

func TestExerciseSubmit_TrimsWhitespace(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthPost(t, server, cookie, "/api/exercises/1/submit", `{"answer":"  hablo  "}`)

	var result struct {
		Correct bool `json:"correct"`
	}
	testutil.MustDecodeJSON(t, resp, &result)

	if !result.Correct {
		t.Error("answer checking should trim whitespace")
	}
}

func TestExerciseSubmit_RecordsProgress(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	// Submit an answer
	resp := testutil.AuthPost(t, server, cookie, "/api/exercises/1/submit", `{"answer":"hablo"}`)
	resp.Body.Close()

	// Check progress summary reflects it
	resp = testutil.AuthGet(t, server, cookie, "/api/progress/summary")
	var summary struct {
		Exercises struct {
			Completed int64 `json:"completed"`
		} `json:"exercises"`
	}
	testutil.MustDecodeJSON(t, resp, &summary)

	if summary.Exercises.Completed < 1 {
		t.Error("exercises completed count should be at least 1 after submitting")
	}
}
