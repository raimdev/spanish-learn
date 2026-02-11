package handler_test

import (
	"testing"

	"spanish-learn-backend/internal/testutil"
)

func TestVocabularyList(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/vocabulary")

	var words []struct {
		ID       int64  `json:"id"`
		Spanish  string `json:"spanish"`
		English  string `json:"english"`
		Category string `json:"category"`
	}
	testutil.MustDecodeJSON(t, resp, &words)

	if len(words) < 50 {
		t.Errorf("expected at least 50 vocabulary words, got %d", len(words))
	}
}

func TestVocabularyList_FilterByCategory(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/vocabulary?category=food")

	var words []struct {
		Category string `json:"category"`
	}
	testutil.MustDecodeJSON(t, resp, &words)

	if len(words) == 0 {
		t.Fatal("should have food vocabulary words")
	}
	for _, w := range words {
		if w.Category != "food" {
			t.Errorf("all words should be in 'food' category, got %q", w.Category)
		}
	}
}

func TestVocabularyCategories(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/vocabulary/categories")

	var categories []string
	testutil.MustDecodeJSON(t, resp, &categories)

	if len(categories) == 0 {
		t.Fatal("should have vocabulary categories")
	}

	// Check that expected categories are present
	expected := map[string]bool{
		"greetings": false,
		"food":      false,
		"family":    false,
		"verbs":     false,
		"colors":    false,
		"numbers":   false,
	}
	for _, c := range categories {
		expected[c] = true
	}
	for cat, found := range expected {
		if !found {
			t.Errorf("expected category %q not found", cat)
		}
	}
}

func TestVocabularyQuiz(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/vocabulary/quiz?count=5")

	var words []struct {
		ID      int64  `json:"id"`
		Spanish string `json:"spanish"`
		English string `json:"english"`
	}
	testutil.MustDecodeJSON(t, resp, &words)

	if len(words) != 5 {
		t.Errorf("expected 5 quiz words, got %d", len(words))
	}

	// Verify each word has content
	for _, w := range words {
		if w.Spanish == "" || w.English == "" {
			t.Error("quiz word should have both Spanish and English")
		}
	}
}

func TestVocabularyQuiz_DefaultCount(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/vocabulary/quiz")

	var words []struct {
		ID int64 `json:"id"`
	}
	testutil.MustDecodeJSON(t, resp, &words)

	if len(words) != 10 {
		t.Errorf("default quiz should return 10 words, got %d", len(words))
	}
}

func TestVocabularyQuiz_WithCategory(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/vocabulary/quiz?count=5&category=greetings")

	var words []struct {
		Category string `json:"category"`
	}
	testutil.MustDecodeJSON(t, resp, &words)

	for _, w := range words {
		if w.Category != "greetings" {
			t.Errorf("quiz with category filter should only return matching words, got %q", w.Category)
		}
	}
}
