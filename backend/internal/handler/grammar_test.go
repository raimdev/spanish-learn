package handler_test

import (
	"net/http"
	"testing"

	"spanish-learn-backend/internal/testutil"
)

func TestGrammarList(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/grammar")

	var lessons []struct {
		ID        int64  `json:"id"`
		Title     string `json:"title"`
		Slug      string `json:"slug"`
		Summary   string `json:"summary"`
		SortOrder int64  `json:"sort_order"`
		Completed bool   `json:"completed"`
	}
	testutil.MustDecodeJSON(t, resp, &lessons)

	if len(lessons) != 5 {
		t.Fatalf("expected 5 grammar lessons, got %d", len(lessons))
	}

	// Check they are sorted
	for i := 1; i < len(lessons); i++ {
		if lessons[i].SortOrder < lessons[i-1].SortOrder {
			t.Error("lessons should be sorted by sort_order")
		}
	}

	// Initially no lessons are completed
	for _, l := range lessons {
		if l.Completed {
			t.Errorf("lesson %q should not be completed initially", l.Title)
		}
	}
}

func TestGrammarGetBySlug(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/grammar/ser-vs-estar")

	var lesson struct {
		ID        int64  `json:"id"`
		Title     string `json:"title"`
		Slug      string `json:"slug"`
		Content   string `json:"content"`
		Completed bool   `json:"completed"`
		Examples  []struct {
			ID      int64  `json:"id"`
			Spanish string `json:"spanish"`
			English string `json:"english"`
		} `json:"examples"`
	}
	testutil.MustDecodeJSON(t, resp, &lesson)

	if lesson.Title != "Ser vs Estar" {
		t.Errorf("expected title 'Ser vs Estar', got %q", lesson.Title)
	}
	if lesson.Content == "" {
		t.Error("lesson content should not be empty")
	}
	if len(lesson.Examples) == 0 {
		t.Error("lesson should have examples")
	}
	if lesson.Completed {
		t.Error("lesson should not be completed initially")
	}
}

func TestGrammarGetBySlug_NotFound(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/grammar/nonexistent-slug")
	testutil.AssertStatus(t, resp, http.StatusNotFound)
	resp.Body.Close()
}

func TestGrammarMarkComplete(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	// Mark lesson 1 as complete
	resp := testutil.AuthPost(t, server, cookie, "/api/grammar/1/complete", "{}")
	testutil.AssertStatus(t, resp, http.StatusOK)
	resp.Body.Close()

	// Verify the lesson is now completed in the list
	resp = testutil.AuthGet(t, server, cookie, "/api/grammar")
	var lessons []struct {
		ID        int64 `json:"id"`
		Completed bool  `json:"completed"`
	}
	testutil.MustDecodeJSON(t, resp, &lessons)

	for _, l := range lessons {
		if l.ID == 1 && !l.Completed {
			t.Error("lesson 1 should be completed after marking")
		}
		if l.ID != 1 && l.Completed {
			t.Errorf("lesson %d should not be completed", l.ID)
		}
	}
}

func TestGrammarMarkComplete_AlsoShowsInDetail(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthPost(t, server, cookie, "/api/grammar/1/complete", "{}")
	resp.Body.Close()

	resp = testutil.AuthGet(t, server, cookie, "/api/grammar/present-tense-regular")
	var lesson struct {
		Completed bool `json:"completed"`
	}
	testutil.MustDecodeJSON(t, resp, &lesson)

	if !lesson.Completed {
		t.Error("lesson detail should show completed=true")
	}
}

func TestGrammarMarkComplete_Idempotent(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	// Mark complete twice — should not error
	resp := testutil.AuthPost(t, server, cookie, "/api/grammar/1/complete", "{}")
	testutil.AssertStatus(t, resp, http.StatusOK)
	resp.Body.Close()

	resp = testutil.AuthPost(t, server, cookie, "/api/grammar/1/complete", "{}")
	testutil.AssertStatus(t, resp, http.StatusOK)
	resp.Body.Close()
}
