-- name: ListGrammarLessons :many
SELECT id, title, slug, summary, sort_order, created_at FROM grammar_lessons ORDER BY sort_order;

-- name: GetGrammarLessonBySlug :one
SELECT * FROM grammar_lessons WHERE slug = ? LIMIT 1;

-- name: GetGrammarExamples :many
SELECT * FROM grammar_examples WHERE lesson_id = ? ORDER BY sort_order;

-- name: GetUserGrammarProgress :one
SELECT * FROM user_grammar_progress WHERE user_id = ? AND lesson_id = ? LIMIT 1;

-- name: UpsertGrammarProgress :exec
INSERT INTO user_grammar_progress (user_id, lesson_id, completed, completed_at)
VALUES (?, ?, 1, datetime('now'))
ON CONFLICT(user_id, lesson_id) DO UPDATE SET completed = 1, completed_at = datetime('now');

-- name: CountCompletedLessons :one
SELECT COUNT(*) FROM user_grammar_progress WHERE user_id = ? AND completed = 1;

-- name: CountTotalLessons :one
SELECT COUNT(*) FROM grammar_lessons;

-- name: ListUserGrammarProgress :many
SELECT lesson_id, completed FROM user_grammar_progress WHERE user_id = ?;
