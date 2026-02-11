package handler_test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"spanish-learn-backend/internal/testutil"
)

func TestLogin_Success(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)

	body := strings.NewReader(`{"username":"test","password":"test"}`)
	resp, err := http.Post(server.URL+"/api/login", "application/json", body)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	testutil.AssertStatus(t, resp, http.StatusOK)

	var found bool
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "session_token" {
			found = true
			if cookie.Value == "" {
				t.Error("session_token cookie should not be empty")
			}
			if !cookie.HttpOnly {
				t.Error("session_token cookie should be HttpOnly")
			}
		}
	}
	if !found {
		t.Error("response should contain session_token cookie")
	}
}

func TestLogin_InvalidCredentials(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)

	body := strings.NewReader(`{"username":"test","password":"wrong"}`)
	resp, err := http.Post(server.URL+"/api/login", "application/json", body)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	testutil.AssertStatus(t, resp, http.StatusUnauthorized)
}

func TestLogin_InvalidBody(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)

	body := strings.NewReader(`not json`)
	resp, err := http.Post(server.URL+"/api/login", "application/json", body)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	testutil.AssertStatus(t, resp, http.StatusBadRequest)
}

func TestMe_Authenticated(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	resp := testutil.AuthGet(t, server, cookie, "/api/me")

	var user struct {
		ID       int64  `json:"id"`
		Username string `json:"username"`
	}
	testutil.MustDecodeJSON(t, resp, &user)

	if user.Username != "test" {
		t.Errorf("expected username 'test', got '%s'", user.Username)
	}
	if user.ID == 0 {
		t.Error("user ID should not be zero")
	}
}

func TestMe_Unauthenticated(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)

	resp, err := http.Get(server.URL + "/api/me")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	testutil.AssertStatus(t, resp, http.StatusUnauthorized)
}

func TestLogout(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)
	cookie := testutil.LoginTestUser(t, server)

	// Logout
	resp := testutil.AuthPost(t, server, cookie, "/api/logout", "{}")
	testutil.AssertStatus(t, resp, http.StatusOK)
	resp.Body.Close()

	// Verify the old cookie no longer works
	resp = testutil.AuthGet(t, server, cookie, "/api/me")
	testutil.AssertStatus(t, resp, http.StatusUnauthorized)
	resp.Body.Close()
}

func TestProtectedRoutes_RequireAuth(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)

	protectedPaths := []string{
		"/api/me",
		"/api/grammar",
		"/api/vocabulary",
		"/api/vocabulary/categories",
		"/api/vocabulary/quiz",
		"/api/exercises",
		"/api/flashcards/decks",
		"/api/flashcards/due",
		"/api/progress/summary",
	}

	for _, path := range protectedPaths {
		resp, err := http.Get(server.URL + path)
		if err != nil {
			t.Fatalf("request to %s failed: %v", path, err)
		}
		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("GET %s: expected 401, got %d", path, resp.StatusCode)
		}
		resp.Body.Close()
	}
}

func TestLogin_ResponseBody(t *testing.T) {
	server, _ := testutil.SetupTestServer(t)

	body := strings.NewReader(`{"username":"test","password":"test"}`)
	resp, err := http.Post(server.URL+"/api/login", "application/json", body)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	if result["message"] != "logged in" {
		t.Errorf("expected message 'logged in', got '%s'", result["message"])
	}
}
