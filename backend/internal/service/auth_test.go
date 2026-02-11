package service

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	db "spanish-learn-backend/db/sqlc"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *db.Queries {
	t.Helper()

	database, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}
	database.Exec("PRAGMA foreign_keys=ON")

	migrationsDir := findMigrationsDir(t)
	runMigrations(t, database, migrationsDir)

	t.Cleanup(func() { database.Close() })
	return db.New(database)
}

func findMigrationsDir(t *testing.T) string {
	t.Helper()
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
		t.Fatalf("failed to read migrations dir: %v", err)
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
			t.Fatalf("failed to execute migration %s: %v", file, err)
		}
	}
}

func TestLogin_ValidCredentials(t *testing.T) {
	queries := setupTestDB(t)
	svc := NewAuthService(queries)

	token, err := svc.Login(context.Background(), "test", "test")
	if err != nil {
		t.Fatalf("login should succeed: %v", err)
	}
	if token == "" {
		t.Fatal("token should not be empty")
	}
	if len(token) != 64 {
		t.Errorf("token should be 64 hex chars, got %d", len(token))
	}
}

func TestLogin_InvalidPassword(t *testing.T) {
	queries := setupTestDB(t)
	svc := NewAuthService(queries)

	_, err := svc.Login(context.Background(), "test", "wrong")
	if err == nil {
		t.Fatal("login should fail with wrong password")
	}
}

func TestLogin_InvalidUsername(t *testing.T) {
	queries := setupTestDB(t)
	svc := NewAuthService(queries)

	_, err := svc.Login(context.Background(), "nonexistent", "test")
	if err == nil {
		t.Fatal("login should fail with nonexistent user")
	}
}

func TestGetUserFromToken_ValidSession(t *testing.T) {
	queries := setupTestDB(t)
	svc := NewAuthService(queries)

	token, err := svc.Login(context.Background(), "test", "test")
	if err != nil {
		t.Fatalf("login failed: %v", err)
	}

	user, err := svc.GetUserFromToken(context.Background(), token)
	if err != nil {
		t.Fatalf("GetUserFromToken should succeed: %v", err)
	}
	if user.Username != "test" {
		t.Errorf("expected username 'test', got '%s'", user.Username)
	}
	if user.ID != 1 {
		t.Errorf("expected user ID 1, got %d", user.ID)
	}
}

func TestGetUserFromToken_InvalidToken(t *testing.T) {
	queries := setupTestDB(t)
	svc := NewAuthService(queries)

	_, err := svc.GetUserFromToken(context.Background(), "invalidtoken")
	if err == nil {
		t.Fatal("should fail with invalid token")
	}
}

func TestLogout_RemovesSession(t *testing.T) {
	queries := setupTestDB(t)
	svc := NewAuthService(queries)

	token, _ := svc.Login(context.Background(), "test", "test")
	svc.Logout(context.Background(), token)

	_, err := svc.GetUserFromToken(context.Background(), token)
	if err == nil {
		t.Fatal("session should be invalid after logout")
	}
}

func TestLogin_CreatesUniqueTokens(t *testing.T) {
	queries := setupTestDB(t)
	svc := NewAuthService(queries)

	token1, _ := svc.Login(context.Background(), "test", "test")
	token2, _ := svc.Login(context.Background(), "test", "test")

	if token1 == token2 {
		t.Error("each login should create a unique token")
	}
}
