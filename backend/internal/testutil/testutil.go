package testutil

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	db "spanish-learn-backend/db/sqlc"
	"spanish-learn-backend/internal/router"

	_ "github.com/mattn/go-sqlite3"
)

// SetupTestDB creates an in-memory SQLite database with all migrations applied.
func SetupTestDB(t *testing.T) (*sql.DB, *db.Queries) {
	t.Helper()

	database, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}

	database.Exec("PRAGMA foreign_keys=ON")

	migrationsDir := findMigrationsDir(t)
	runMigrations(t, database, migrationsDir)

	queries := db.New(database)
	t.Cleanup(func() { database.Close() })

	return database, queries
}

// SetupTestServer creates a test HTTP server with the full router.
func SetupTestServer(t *testing.T) (*httptest.Server, *db.Queries) {
	t.Helper()

	_, queries := SetupTestDB(t)
	r := router.New(queries, "http://localhost:3000")
	server := httptest.NewServer(r)
	t.Cleanup(server.Close)

	return server, queries
}

// LoginTestUser logs in the test user and returns the session cookie.
func LoginTestUser(t *testing.T, server *httptest.Server) *http.Cookie {
	t.Helper()

	body := strings.NewReader(`{"username":"test","password":"test"}`)
	resp, err := http.Post(server.URL+"/api/login", "application/json", body)
	if err != nil {
		t.Fatalf("login request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("login returned status %d", resp.StatusCode)
	}

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "session_token" {
			return cookie
		}
	}
	t.Fatal("no session_token cookie in login response")
	return nil
}

// AuthGet makes an authenticated GET request.
func AuthGet(t *testing.T, server *httptest.Server, cookie *http.Cookie, path string) *http.Response {
	t.Helper()
	req, _ := http.NewRequest("GET", server.URL+path, nil)
	req.AddCookie(cookie)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("GET %s failed: %v", path, err)
	}
	return resp
}

// AuthPost makes an authenticated POST request with JSON body.
func AuthPost(t *testing.T, server *httptest.Server, cookie *http.Cookie, path string, jsonBody string) *http.Response {
	t.Helper()
	req, _ := http.NewRequest("POST", server.URL+path, strings.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(cookie)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("POST %s failed: %v", path, err)
	}
	return resp
}

// DecodeJSON decodes the response body into the given target.
func DecodeJSON(t *testing.T, resp *http.Response, target interface{}) {
	t.Helper()
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}
}

func findMigrationsDir(t *testing.T) string {
	t.Helper()
	// Walk up from the working directory to find db/migrations
	dir, _ := os.Getwd()
	for i := 0; i < 10; i++ {
		candidate := filepath.Join(dir, "db", "migrations")
		if info, err := os.Stat(candidate); err == nil && info.IsDir() {
			return candidate
		}
		dir = filepath.Dir(dir)
	}
	t.Fatal("could not find db/migrations directory")
	return ""
}

func runMigrations(t *testing.T, database *sql.DB, dir string) {
	t.Helper()

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("failed to read migrations directory: %v", err)
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".sql") {
			files = append(files, entry.Name())
		}
	}
	sort.Strings(files)

	for _, file := range files {
		content, err := os.ReadFile(filepath.Join(dir, file))
		if err != nil {
			t.Fatalf("failed to read migration %s: %v", file, err)
		}
		if _, err := database.Exec(string(content)); err != nil {
			t.Fatalf("failed to execute migration %s: %v\nSQL: %s", file, err, string(content)[:min(200, len(content))])
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// AssertStatus checks the response status code.
func AssertStatus(t *testing.T, resp *http.Response, expected int) {
	t.Helper()
	if resp.StatusCode != expected {
		t.Errorf("expected status %d, got %d", expected, resp.StatusCode)
	}
}

// MustDecodeJSON decodes the response and asserts a 200 status.
func MustDecodeJSON(t *testing.T, resp *http.Response, target interface{}) {
	t.Helper()
	if resp.StatusCode != http.StatusOK {
		var errBody map[string]string
		json.NewDecoder(resp.Body).Decode(&errBody)
		resp.Body.Close()
		t.Fatalf("expected status 200, got %d: %v", resp.StatusCode, errBody)
	}
	DecodeJSON(t, resp, target)
}

// FormatJSON returns a pretty-printed JSON string for debugging.
func FormatJSON(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return fmt.Sprintf("%s", b)
}
